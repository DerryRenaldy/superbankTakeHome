export interface BankAccount {
  account_number: string;
  balance: number;
  currency: string;
}

export interface Pocket {
  name: string;
  balance: number;
  currency: string;
}

export interface TermDeposit {
  amount: number;
  currency: string;
  interest_rate: number;
  maturity_date: string;
}

export interface Customer {
  name: string;
  bank_accounts: BankAccount[];
  pockets: Pocket[];
  term_deposits: TermDeposit[];
}
