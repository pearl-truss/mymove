/* global cy */
import { milmoveAppName } from '../../support/constants';

describe('allows a SM to request a payment', function () {
  const moveID = 'f9f10492-587e-43b3-af2a-9f67d2ac8757';
  beforeEach(() => {
    cy.removeFetch();
    cy.server();
    cy.route('POST', '**/internal/uploads').as('postUploadDocument');
    cy.route('POST', '**/moves/**/weight_ticket').as('postWeightTicket');
    cy.route('POST', '**/moves/**/moving_expense_documents').as('postMovingExpense');
    cy.route('POST', '**/internal/personally_procured_move/**/request_payment').as('requestPayment');
    cy.route('POST', '**/moves/**/signed_certifications').as('signedCertifications');
  });

  it('service member goes through entire request payment flow', () => {
    cy.signInAsUserPostRequest(milmoveAppName, '8e0d7e98-134e-4b28-bdd1-7d6b1ff34f9e');
    serviceMemberStartsPPMPaymentRequest();
    serviceMemberSubmitsWeightTicket('CAR', true);
    serviceMemberChecksNumberOfWeightTickets('2nd');
    serviceMemberSubmitsWeightTicket('BOX_TRUCK', false);
    serviceMemberViewsExpensesLandingPage();
    serviceMemberUploadsExpenses(false);
    serviceMemberReviewsDocuments();
    serviceMemberEditsPaymentRequest();
    serviceMemberAddsWeightTicketSetWithMissingDocuments();
    serviceMemberSubmitsPaymentRequestWithMissingDocuments();
  });

  it('service member reads introduction to ppm payment and goes back to homepage', () => {
    cy.signInAsUserPostRequest(milmoveAppName, '745e0eba-4028-4c78-a262-818b00802748');
    serviceMemberStartsPPMPaymentRequest();
  });

  it('service member can save a weight ticket for later', () => {
    cy.signInAsUserPostRequest(milmoveAppName, '745e0eba-4028-4c78-a262-818b00802748');

    cy.visit(`/moves/${moveID}/ppm-weight-ticket`);
    serviceMemberCanFinishWeightTicketLater('BOX_TRUCK');
  });

  it('service member submits weight tickets without any documents', () => {
    cy.signInAsUserPostRequest(milmoveAppName, '745e0eba-4028-4c78-a262-818b00802748');
    cy.visit(`/moves/${moveID}/ppm-weight-ticket`);
    serviceMemberSubmitsWeightsTicketsWithoutReceipts();
  });

  it('service member requests a car + trailer weight ticket payment', () => {
    cy.signInAsUserPostRequest(milmoveAppName, '745e0eba-4028-4c78-a262-818b00802748');
    cy.visit(`/moves/${moveID}/ppm-weight-ticket`);
    serviceMemberSubmitsCarTrailerWeightTicket();
  });

  it('service member starting at review page returns to review page after adding a weight ticket', () => {
    cy.signInAsUserPostRequest(milmoveAppName, '745e0eba-4028-4c78-a262-818b00802748');
    cy.visit(`/moves/${moveID}/ppm-payment-review`);
    cy.location().should((loc) => {
      expect(loc.pathname).to.match(/^\/moves\/[^/]+\/ppm-payment-review/);
    });
    cy.get('[data-cy="weight-ticket-link"]').click();
    cy.location().should((loc) => {
      expect(loc.pathname).to.match(/^\/moves\/[^/]+\/ppm-weight-ticket/);
    });
    serviceMemberSubmitsWeightTicket('CAR', false);
    cy.location().should((loc) => {
      expect(loc.pathname).to.match(/^\/moves\/[^/]+\/ppm-payment-review/);
    });
  });

  it('service member starting at review page returns to review page after adding an expense', () => {
    cy.signInAsUserPostRequest(milmoveAppName, '745e0eba-4028-4c78-a262-818b00802748');
    cy.visit(`/moves/${moveID}/ppm-payment-review`);
    cy.location().should((loc) => {
      expect(loc.pathname).to.match(/^\/moves\/[^/]+\/ppm-payment-review/);
    });
    cy.get('[data-cy="expense-link"]').click();
    cy.location().should((loc) => {
      expect(loc.pathname).to.match(/^\/moves\/[^/]+\/ppm-expenses/);
    });
    serviceMemberUploadsExpenses(false);
    cy.location().should((loc) => {
      expect(loc.pathname).to.match(/^\/moves\/[^/]+\/ppm-payment-review/);
    });
  });

  it('service member can skip weight tickets and expenses if already have one', () => {
    cy.signInAsUserPostRequest(milmoveAppName, '745e0eba-4028-4c78-a262-818b00802748');
    cy.visit(`/moves/${moveID}/ppm-weight-ticket`);
    serviceMemberSubmitsWeightTicket('CAR', true);
    serviceMemberSkipsStep();
    serviceMemberViewsExpensesLandingPage();
    serviceMemberUploadsExpenses();
    serviceMemberSkipsStep();
  });

  it('service member with old weight tickets can see and delete them', () => {
    cy.signInAsUserPostRequest(milmoveAppName, 'beccca28-6e15-40cc-8692-261cae0d4b14');
    cy.get('[data-cy="edit-payment-request"]').contains('Edit Payment Request').should('exist').click();
    cy.get('.ticket-item').first().should('not.contain', 'set');
    cy.get('[data-cy="delete-ticket"]').first().click();
    cy.get('[data-cy="delete-confirmation-button"]').click();
    cy.get('.ticket-item').should('not.exist');
  });
});

function serviceMemberSkipsStep() {
  cy.get('[data-cy=skip]').contains('Skip').click();
}

function serviceMemberSubmitsPaymentRequestWithMissingDocuments() {
  cy.get('.review-customer-agreement a').contains('Legal Agreement').click();
  cy.location().should((loc) => {
    expect(loc.pathname).to.match(/^\/ppm-customer-agreement/);
  });
  cy.get('.usa-button').contains('Back').click();
  cy.location().should((loc) => {
    expect(loc.pathname).to.match(/^\/moves\/[^/]+\/ppm-payment-review/);
  });
  cy.get('.missing-label').contains('Your estimated payment is unknown');

  cy.get('input[id="agree-checkbox"]').check({ force: true });

  cy.get('button').contains('Submit Request').should('be.enabled').click();
  cy.wait('@signedCertifications');
  cy.wait('@requestPayment');

  cy.get('.usa-alert--warning').contains('Payment request is missing info');
  cy.get('.usa-alert--warning').contains(
    'You will need to contact your local PPPO office to resolve your missing weight ticket.',
  );

  cy.get('.title').contains('Next step: Contact the PPPO office');
  cy.get('.missing-label').contains('Unknown');
}

function serviceMemberReviewsDocuments() {
  cy.location().should((loc) => {
    expect(loc.pathname).to.match(/^\/moves\/[^/]+\/ppm-payment-review/);
  });
  cy.get('.review-customer-agreement a').contains('Legal Agreement').click();
  cy.location().should((loc) => {
    expect(loc.pathname).to.match(/^\/ppm-customer-agreement/);
  });
  cy.get('[data-cy="back-button"]').contains('Back').click();
  cy.location().should((loc) => {
    expect(loc.pathname).to.match(/^\/moves\/[^/]+\/ppm-payment-review/);
  });
  cy.get('input[id="agree-checkbox"]').check({ force: true });
  cy.contains(`You're requesting a payment of $`);
  cy.get('button').contains('Submit Request').should('be.enabled').click();
  cy.wait('@signedCertifications');
  cy.wait('@requestPayment');
}
function serviceMemberDeletesDocuments() {
  cy.location().should((loc) => {
    expect(loc.pathname).to.match(/^\/moves\/[^/]+\/ppm-payment-review/);
  });
  cy.get('.ticket-item').should('have.length', 4);
  cy.get('[data-cy="delete-ticket"]').first().click();
  cy.get('[data-cy="delete-confirmation-button"]').click();
  cy.get('.ticket-item').should('have.length', 3);
}
function serviceMemberEditsPaymentRequest() {
  cy.get('.usa-alert--success').contains('Payment request submitted').should('exist');
  cy.get('[data-cy="edit-payment-request"]').contains('Edit Payment Request').should('exist').click();
  cy.get('[data-cy=weight-ticket-link]').should('exist').click();
  serviceMemberSubmitsWeightTicket('CAR', false);
  serviceMemberDeletesDocuments();
  serviceMemberReviewsDocuments();
}
function serviceMemberAddsWeightTicketSetWithMissingDocuments(hasAnother = false) {
  cy.get('[data-cy="edit-payment-request"]').contains('Edit Payment Request').should('exist').click();
  cy.get('[data-cy=weight-ticket-link]').should('exist').click();

  cy.get('select[name="weight_ticket_set_type"]').select('BOX_TRUCK');

  cy.get('input[name="vehicle_nickname"]').type('Nickname');

  cy.get('input[name="empty_weight"]').type('1000');
  cy.get('input[name="missingEmptyWeightTicket"]').check({ force: true });

  cy.get('input[name="full_weight"]').type('5000');
  cy.upload_file('[data-cy=full-weight-upload] .filepond--root', 'top-secret.png');
  cy.wait('@postUploadDocument');

  cy.get('input[name="weight_ticket_date"]').type('6/2/2018{enter}').blur();
  cy.get('input[name="additional_weight_ticket"][value="Yes"]').should('not.be.checked');
  cy.get('input[name="additional_weight_ticket"][value="No"]').should('be.checked');

  if (hasAnother) {
    cy.get('input[name="additional_weight_ticket"][value="Yes"]+label').click();
    cy.get('input[name="additional_weight_ticket"][value="Yes"]').should('be.checked');
    cy.get('button').contains('Save & Add Another').click();
    cy.wait('@postWeightTicket').its('status').should('eq', 200);
    cy.get('[data-cy=documents-uploaded]').should('exist');
  } else {
    cy.get('button').contains('Save & Continue').click();
    cy.wait('@postWeightTicket').its('status').should('eq', 200);
  }
}
function serviceMemberViewsExpensesLandingPage() {
  cy.location().should((loc) => {
    expect(loc.pathname).to.match(/^\/moves\/[^/]+\/ppm-expenses-intro/);
  });

  cy.get('[data-cy=documents-uploaded]').should('exist');
  cy.get('button').contains('Continue').should('be.disabled');

  cy.get('[type="radio"]').first().should('be.not.checked');
  cy.get('[type="radio"]').last().should('be.not.checked');

  cy.get('a')
    .contains('More about expenses')
    .should('have.attr', 'href')
    .and('match', /^\/allowable-expenses/);

  cy.get('input[name="hasExpenses"][value="Yes"]').should('not.be.checked');
  cy.get('input[name="hasExpenses"][value="No"]').should('not.be.checked');
  cy.get('input[name="hasExpenses"][value="Yes"]+label').click();

  cy.get('button').contains('Continue').should('be.enabled').click();
}

function serviceMemberUploadsExpenses(hasAnother = true, expenseNumber = null) {
  cy.location().should((loc) => {
    expect(loc.pathname).to.match(/^\/moves\/[^/]+\/ppm-expenses/);
  });

  if (expenseNumber) {
    cy.contains(`Expense ${expenseNumber}`);
  }
  cy.get('[data-cy=documents-uploaded]').should('exist');

  cy.get('select[name="moving_expense_type"]').select('GAS');
  cy.get('input[name="title"]').type('title');
  cy.get('input[name="requested_amount_cents"]').type('1000');

  cy.upload_file('.filepond--root:first', 'top-secret.png');
  cy.wait('@postUploadDocument');
  cy.get('[data-filepond-item-state="processing-complete"]').should('have.length', 1);

  cy.get('input[name="missingReceipt"]').should('not.be.checked');
  cy.get('input[name="paymentMethod"][value="GTCC"]').should('be.checked');
  cy.get('input[name="paymentMethod"][value="OTHER"]').should('not.be.checked');
  cy.get('input[name="haveMoreExpenses"][value="Yes"]').should('not.be.checked');
  cy.get('input[name="haveMoreExpenses"][value="No"]').should('be.checked');
  cy.get('input[name="haveMoreExpenses"][value="Yes"]+label').click();
  if (hasAnother) {
    cy.get('button').contains('Save & Add Another').click();
    cy.wait('@postMovingExpense').its('status').should('eq', 200);
    cy.get('[data-cy=documents-uploaded]').should('exist');
  } else {
    cy.get('input[name="haveMoreExpenses"][value="No"]+label').click();
    cy.get('input[name="haveMoreExpenses"][value="No"]').should('be.checked');
    cy.get('button').contains('Save & Continue').click();
    cy.wait('@postMovingExpense').its('status').should('eq', 200);
  }
}

function serviceMemberSubmitsCarTrailerWeightTicket() {
  cy.get('select[name="weight_ticket_set_type"]').select('CAR_TRAILER');

  cy.get('input[name="vehicle_make"]').type('Make');
  cy.get('input[name="vehicle_model"]').type('Model');

  cy.contains('Do you own this trailer').children('a').should('have.attr', 'href', '/trailer-criteria');

  cy.get('input[name="isValidTrailer"][value="Yes"]').should('not.be.checked');
  cy.get('input[name="isValidTrailer"][value="No"]').should('be.checked');
  cy.get('input[name="isValidTrailer"][value="Yes"]+label').click();

  cy.upload_file('[data-cy=trailer-upload] .filepond--root', 'top-secret.png');
  cy.wait('@postUploadDocument');
  cy.get('[data-filepond-item-state="processing-complete"]').should('have.length', 1);

  cy.get('input[name="missingDocumentation"]').check({ force: true });

  cy.get('input[name="empty_weight"]').type('1000');
  cy.get('input[name="missingEmptyWeightTicket"]').check({ force: true });

  cy.get('input[name="full_weight"]').type('5000');
  cy.upload_file('.filepond--root:last', 'top-secret.png');
  cy.wait('@postUploadDocument');
  cy.get('[data-filepond-item-state="processing-complete"]').should('have.length', 2);
  cy.get('input[name="weight_ticket_date"]').type('6/2/2018{enter}').blur();

  cy.get('input[name="additional_weight_ticket"][value="Yes"]').should('not.be.checked');
  cy.get('input[name="additional_weight_ticket"][value="No"]').should('be.checked');
}
function serviceMemberCanFinishWeightTicketLater(vehicleType) {
  cy.get('select[name="weight_ticket_set_type"]').select(vehicleType);

  if (vehicleType === 'BOX_TRUCK' || vehicleType === 'PRO_GEAR') {
    cy.get('input[name="vehicle_nickname"]').type('Nickname');
  } else if (vehicleType === 'CAR' || vehicleType === 'CAR_TRAILER') {
    cy.get('input[name="vehicle_make"]').type('Make');
    cy.get('input[name="vehicle_model"]').type('Model');
  }

  cy.get('input[name="empty_weight"]').type('1000');
  cy.upload_file('[data-cy=empty-weight-upload] .filepond--root', 'top-secret.png');
  cy.wait('@postUploadDocument');
  cy.get('[data-filepond-item-state="processing-complete"]').should('have.length', 1);

  cy.get('input[name="full_weight"]').type('5000');
  cy.upload_file('[data-cy=full-weight-upload] .filepond--root', 'top-secret.png');
  cy.wait('@postUploadDocument');
  cy.get('[data-filepond-item-state="processing-complete"]').should('have.length', 2);
  cy.get('input[name="weight_ticket_date"]').type('6/2/2018{enter}').blur();

  cy.get('button').contains('Finish Later').click();

  cy.get('button').contains('Cancel').click();

  cy.location().should((loc) => {
    expect(loc.pathname).to.match(/^\/moves\/[^/]+\/ppm-weight-ticket/);
  });

  cy.get('button').contains('Finish Later').click();

  cy.get('button').contains('OK').click();

  cy.location().should((loc) => {
    expect(loc.pathname).to.match(/^\/$/);
  });
}

function serviceMemberSubmitsWeightsTicketsWithoutReceipts() {
  cy.get('select[name="weight_ticket_set_type"]').select('CAR_TRAILER');
  cy.get('input[name="vehicle_make"]').type('Make');
  cy.get('input[name="vehicle_model"]').type('Model');
  cy.get('input[name="empty_weight"]').type('1000');
  cy.get('input[name="full_weight"]').type('2000');
  cy.upload_file('[data-cy=full-weight-upload] .filepond--root', 'top-secret.png');
  cy.wait('@postUploadDocument');
  cy.get('input[name="isValidTrailer"][value="Yes"]+label').click();
  cy.get('input[name="missingDocumentation"]+label').click();
  cy.get('[data-cy=trailer-warning]').contains(
    'If your state does not provide a registration or bill of sale for your trailer, you may write and upload a signed and dated statement certifying that you or your spouse own the trailer and meets the trailer criteria. Upload your statement using the proof of ownership field.',
  );
  cy.get('input[name="missingDocumentation"]+label').click({ force: false });
  cy.upload_file('[data-cy=trailer-upload] .filepond--root', 'top-secret.png');
  cy.wait('@postUploadDocument');

  cy.get('input[name="missingEmptyWeightTicket"]+label').click();
  cy.get('[data-cy=empty-warning]').contains(
    'Contact your local Transportation Office (PPPO) to let them know you’re missing this weight ticket. For now, keep going and enter the info you do have.',
  );

  cy.get('input[name="weight_ticket_date"]').type('6/2/2018{enter}').blur();

  cy.get('input[name="additional_weight_ticket"][value="Yes"]+label').click();
  cy.get('input[name="additional_weight_ticket"][value="Yes"]').should('be.checked');
  cy.get('button').contains('Save & Add Another').click();
  cy.wait('@postWeightTicket');
}

function serviceMemberStartsPPMPaymentRequest() {
  cy.contains('Request Payment').click();
  cy.get('input[name="actual_move_date"]').type('6/20/2018{enter}').blur();
  cy.get('button').contains('Get Started').click();
}

function serviceMemberChecksNumberOfWeightTickets(ordinal) {
  cy.contains(`Weight Tickets - ${ordinal} set`);
  cy.get('[data-cy=documents-uploaded]').should('exist');
}

function serviceMemberSubmitsWeightTicket(vehicleType, hasAnother = true) {
  cy.get('select[name="weight_ticket_set_type"]').select(vehicleType);

  if (vehicleType === 'BOX_TRUCK' || vehicleType === 'PRO_GEAR') {
    cy.get('input[name="vehicle_nickname"]').type('Nickname');
  } else if (vehicleType === 'CAR' || vehicleType === 'CAR_TRAILER') {
    cy.get('input[name="vehicle_make"]').type('Make');
    cy.get('input[name="vehicle_model"]').type('Model');
  }

  cy.get('input[name="empty_weight"]').type('1000');

  cy.upload_file('[data-cy=empty-weight-upload] .filepond--root', 'top-secret.png');
  cy.wait('@postUploadDocument');
  cy.get('[data-filepond-item-state="processing-complete"]').should('have.length', 1);

  cy.get('input[name="full_weight"]').type('5000');
  cy.upload_file('[data-cy=full-weight-upload] .filepond--root', 'top-secret.png');
  cy.wait('@postUploadDocument');
  cy.get('[data-filepond-item-state="processing-complete"]').should('have.length', 2);
  cy.get('input[name="weight_ticket_date"]').type('6/2/2018{enter}').blur();
  cy.get('input[name="additional_weight_ticket"][value="Yes"]').should('not.be.checked');
  cy.get('input[name="additional_weight_ticket"][value="No"]').should('be.checked');
  if (hasAnother) {
    cy.get('input[name="additional_weight_ticket"][value="Yes"]+label').click();
    cy.get('input[name="additional_weight_ticket"][value="Yes"]').should('be.checked');
    cy.get('button').contains('Save & Add Another').click();
    cy.wait('@postWeightTicket').its('status').should('eq', 200);
    cy.get('[data-cy=documents-uploaded]').should('exist');
  } else {
    cy.get('button').contains('Save & Continue').click();
    cy.wait('@postWeightTicket').its('status').should('eq', 200);
  }
}
