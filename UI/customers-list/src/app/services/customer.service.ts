import {Injectable} from '@angular/core';
import {HttpClient} from "@angular/common/http";
import {Observable} from "rxjs";
import {CustomersResponse} from "../data/customers.response";

@Injectable({
  providedIn: 'root'
})
export class CustomerService {

  constructor(private httpClient: HttpClient) {
  }

  public listCustomers(): Observable<Array<CustomersResponse>> {
    return this.httpClient.get<Array<CustomersResponse>>("http://localhost:8080/customers");
  }


}
