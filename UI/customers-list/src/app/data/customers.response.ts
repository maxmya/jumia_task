import {Country} from "./country.model";
import {Customer} from "./customer.model";

export interface CustomersResponse {
  country: Country;
  customers: Array<Customer>;
}
