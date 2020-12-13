import { useEffect, useState } from 'react'
import Link from 'next/link'
import { getSalaryFromApiByUuid } from '../../consumers/api'
import { formatMoney } from '../Pages/SalaryReportPage'
import {
  DescriptionOrUuid,
  DescriptionOrUuidFn,
} from '../Atoms/DescriptionOrUuid'

export function SalaryCard({ salary: salaryUuid }) {
  const [salary, setSalary] = useState({})
  useEffect(() => {
    getSalaryFromApiByUuid(salaryUuid).then(setSalary)
  }, [salaryUuid])

  const removeSalary = () => {
    if (confirm('Delete salary ' + DescriptionOrUuidFn(salary) + '?')) {
      localStorage.setItem(
        'mySalaries',
        JSON.stringify(
          JSON.parse(localStorage.getItem('mySalaries')).filter(
            (uuid) => uuid !== salaryUuid
          )
        )
      )
      setSalary({})
    }
  }
  if (!salary?.id) return null
  return (
    <div className="col col-md-4 col-sm-12 mb-4">
      <div className="card ">
        <div className="card-header bg-dark">
          <h3 className="card-title text-truncate text-white mb-0">
            <DescriptionOrUuid salary={salary} />
          </h3>
        </div>
        <div className="card-body">
          <table className="table table-bordered">
            <tbody>
              <tr>
                <td>Amount</td>
                <td>
                  {formatMoney(salary.amount, salary.currency_exchange.from)}
                </td>
              </tr>
              <tr>
                <td>From currency</td>
                <td>{salary.currency_exchange.from}</td>
              </tr>
              <tr>
                <td>To currency</td>
                <td>{salary.currency_exchange.to}</td>
              </tr>
              <tr>
                <td>Rate tax</td>
                <td>{salary.currency_exchange.rate_tax.toFixed(2)}</td>
              </tr>
              <tr>
                <td>Fixed tax amount</td>
                <td>
                  {formatMoney(
                    salary.currency_exchange.fixed_tax_amount,
                    salary.currency_exchange.to
                  )}
                </td>
              </tr>
            </tbody>
          </table>
          <Link href={'/salary/' + salaryUuid}>
            <button className="btn btn-success btn-block mt-4">
              Access report
            </button>
          </Link>

          <button
            className="btn btn-outline-danger btn-block btn-sm mt-2"
            onClick={removeSalary}
          >
            Remove
          </button>
        </div>
      </div>
    </div>
  )
}
