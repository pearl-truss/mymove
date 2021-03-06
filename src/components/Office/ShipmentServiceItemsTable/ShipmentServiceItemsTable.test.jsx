import React from 'react';
import { shallow } from 'enzyme';

import ShipmentServiceItemsTable from './ShipmentServiceItemsTable';

describe('Shipment Service Items Table', () => {
  it('renders the container successfully', () => {
    const wrapper = shallow(<ShipmentServiceItemsTable shipmentType="hhg" />);
    expect(wrapper.find('table.serviceItemsTable').exists()).toBe(true);
    expect(wrapper.find('h4').text()).toEqual('Service items for this shipment 6 items');
    expect(wrapper.find('th').text()).toEqual('Service item');
    expect(wrapper.find('td').length).toBe(6);
    expect(wrapper.find('td').at(0).text()).toEqual('Domestic linehaul');
    expect(wrapper.find('td').at(1).text()).toEqual('Fuel surcharge');
    expect(wrapper.find('td').at(2).text()).toEqual('Domestic origin price');
    expect(wrapper.find('td').at(3).text()).toEqual('Domestic destination price');
    expect(wrapper.find('td').at(4).text()).toEqual('Domestic packing');
    expect(wrapper.find('td').at(5).text()).toEqual('Domestic unpacking');
  });

  it('renders the nts shipment type with correct items', () => {
    const wrapper = shallow(<ShipmentServiceItemsTable shipmentType="nts" />);
    expect(wrapper.find('table.serviceItemsTable').exists()).toBe(true);
    expect(wrapper.find('h4').text()).toEqual('Service items for this shipment 5 items');
    expect(wrapper.find('th').text()).toEqual('Service item');
    expect(wrapper.find('td').length).toBe(5);
    expect(wrapper.find('td').at(0).text()).toEqual('Domestic linehaul');
    expect(wrapper.find('td').at(1).text()).toEqual('Fuel surcharge');
    expect(wrapper.find('td').at(2).text()).toEqual('Domestic origin price');
    expect(wrapper.find('td').at(3).text()).toEqual('Domestic destination price');
    expect(wrapper.find('td').at(4).text()).toEqual('Domestic unpacking');
  });
});
