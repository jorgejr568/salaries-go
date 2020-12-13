import { useState } from 'react'
import { createSalaryOnApi } from '../../consumers/api'
import { useRouter } from 'next/router'

export function SalaryCreatePage() {
  const router = useRouter()
  const [loading, setLoading] = useState(false)
  const [description, setDescription] = useState('')
  const [amount, setAmount] = useState(0.0)
  const [currencyTo, setCurrencyTo] = useState('')
  const [currencyFrom, setCurrencyFrom] = useState('')
  const [currencyTax, setCurrencyTax] = useState(0.0)
  const [currencyFixedTax, setCurrencyFixedTax] = useState(0)

  const parseMoney = (v) => parseFloat(v.toString().replace(/,/g, '.'))
  const parseAmount = (v) => parseInt(parseMoney(v) * 100)
  const saveSalary = () => {
    setLoading(true)
    createSalaryOnApi({
      amount: parseAmount(amount),
      description,
      currency_exchange: {
        from: currencyFrom,
        to: currencyTo,
        rate_tax: parseMoney(currencyTax),
        fixed_tax_amount: parseAmount(currencyFixedTax),
      },
    })
      .then((salary) => {
        let lsSalaries = JSON.parse(localStorage.getItem('mySalaries'))
        if (!Array.isArray(lsSalaries)) {
          lsSalaries = []
        }
        localStorage.setItem(
          'mySalaries',
          JSON.stringify([...lsSalaries, salary.uuid])
        )

        router.push('/')
      })
      .finally(() => setLoading(false))
  }

  const handleInput = (fn) => (e) => {
    fn(e.target.value)
  }

  return (
    <>
      <div className="container">
        <div className="row d-flex justify-content-center">
          <div className="col-md-6 col-sm-12">
            <h3>Attach salary</h3>
            <hr />
            <form onSubmit={(e) => e.preventDefault() & saveSalary()}>
              <input
                type="text"
                className="form-control"
                placeholder="Description"
                onChange={handleInput(setDescription)}
              />

              <input
                type="number"
                step="0.01"
                className="form-control mt-2"
                placeholder="Amount"
                onChange={handleInput(setAmount)}
              />

              <input
                type="text"
                maxLength={3}
                className="form-control mt-2"
                placeholder="From currency"
                onChange={handleInput(setCurrencyFrom)}
              />

              <input
                type="text"
                maxLength={3}
                className="form-control mt-2"
                placeholder="To currency"
                onChange={handleInput(setCurrencyTo)}
              />

              <input
                type="number"
                step="0.01"
                className="form-control mt-2"
                placeholder="Rate tax (amount per Currency from)"
                onChange={handleInput(setCurrencyTax)}
              />

              <input
                type="number"
                step="0.01"
                className="form-control mt-2"
                placeholder="Fixed tax (fixed amount that will be taxed every transaction on Currency to)"
                onChange={handleInput(setCurrencyFixedTax)}
              />

              <button className="btn btn-success btn-block mt-2" type="submit">
                {loading ? 'Saving...' : 'Save'}
              </button>
            </form>
          </div>
        </div>
      </div>
    </>
  )
}
