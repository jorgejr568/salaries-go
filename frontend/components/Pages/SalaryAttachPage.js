import { useState } from 'react'
import { getSalaryFromApiByUuid } from '../../consumers/api'
import { useRouter } from 'next/router'

export function SalaryAttachPage() {
  const [loading, setLoading] = useState(false)
  const [uuid, setUuid] = useState('')
  const router = useRouter()

  const attachSalary = () => {
    if (!uuid.length) return
    setLoading(true)
    getSalaryFromApiByUuid(uuid)
      .then(() => {
        let lsSalaries = JSON.parse(localStorage.getItem('mySalaries'))
        if (!Array.isArray(lsSalaries)) {
          lsSalaries = []
        }

        localStorage.setItem(
          'mySalaries',
          JSON.stringify([...lsSalaries, uuid])
        )
        router.push('/')
      })
      .finally(() => setLoading(false))
  }

  return (
    <>
      <div className="container">
        <div className="row d-flex justify-content-center">
          <div className="col-md-6 col-sm-12">
            <h3>Attach salary</h3>
            <hr />
            <form onSubmit={(e) => e.preventDefault() & attachSalary()}>
              <input
                type="text"
                className="form-control"
                placeholder="UUID"
                onChange={(e) => setUuid(e.target.value?.trim())}
              />

              <button className="btn btn-success btn-block mt-2" type="submit">
                {loading ? 'Attaching' : 'Attach'}
              </button>
            </form>
          </div>
        </div>
      </div>
    </>
  )
}
