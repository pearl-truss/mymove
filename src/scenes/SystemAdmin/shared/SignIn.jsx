import React from 'react';
import qs from 'query-string';

import { withContext } from 'shared/AppContext';
import Alert from 'shared/Alert';
import styles from './SignIn.module.scss';

const SignIn = ({ context, location }) => {
  const error = qs.parse(location.search).error;
  return (
    <div>
      <div>&nbsp;</div>
      <div>
        {error && (
          <div>
            <Alert type="error" heading="An error occurred">
              There was an error during your last sign in attempt. Please try again.
              <br />
              Error code: {error}
            </Alert>
            <br />
          </div>
        )}
        <h2 className="align-center">Welcome to {context.siteName}!</h2>
        <br />
        <p className="align-center">
          This is a new system from USTRANSCOM to support the relocation of families during PCS.
        </p>
        <div className="align-center">
          <a href="/auth/login-gov" className={styles['usa-button']}>
            Sign in
          </a>
        </div>
      </div>
    </div>
  );
};

export default withContext(SignIn);
