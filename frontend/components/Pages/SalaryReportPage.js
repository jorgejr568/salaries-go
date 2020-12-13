import Link from 'next/link'
import { DescriptionOrUuid } from '../Atoms/DescriptionOrUuid'

export const formatMoney = (amount, currency) => {
  return new Intl.NumberFormat('en', { style: 'currency', currency }).format(
    amount / 100
  )
}

export function SalaryReportPage({ salaryExchange }) {
  const { salary, amounts } = salaryExchange
  return (
    <>
      <div className="container">
        <div className="row d-flex justify-content-center">
          <div className="col-md-4 col-sm-12">
            <h3 className="text-truncate">
              Salary <DescriptionOrUuid salary={salary} />
            </h3>
            <hr />
            <ul className="list-group">
              <li className="list-group-item d-flex">
                Initial amount:
                <strong className="ml-auto">
                  {formatMoney(salary.amount, salary.currency_exchange.from)}
                </strong>
              </li>
              <li className="list-group-item d-flex">
                Exchanged amount:
                <strong className="ml-auto">
                  {formatMoney(amounts.total, salary.currency_exchange.to)}
                </strong>
              </li>
              <li className="list-group-item d-flex">
                Taxed amount:
                <strong className="ml-auto text-danger">
                  -
                  {formatMoney(
                    amounts.total - amounts.after_taxes,
                    salary.currency_exchange.to
                  )}
                </strong>
              </li>
              <li className="list-group-item d-flex">
                Final amount:
                <strong className="ml-auto">
                  {formatMoney(
                    amounts.after_taxes,
                    salary.currency_exchange.to
                  )}
                </strong>
              </li>
            </ul>
            <Link href={'/'}>
              <button className="btn btn-danger btn-block mt-4">Back</button>
            </Link>
          </div>
        </div>
      </div>
    </>
  )
}
