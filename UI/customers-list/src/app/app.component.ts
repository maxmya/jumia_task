import {Component} from '@angular/core';
import {Observable} from "rxjs";
import {CustomersResponse} from "./data/customers.response";
import {CustomerService} from "./services/customer.service";

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent {

  customers$: Observable<Array<CustomersResponse>>

  constructor(private customerService: CustomerService) {
    this.customers$ = customerService.listCustomers();
  }


}
