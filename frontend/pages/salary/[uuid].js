import { useState, useEffect } from 'react'
import { useRouter } from 'next/router'
import { Layout } from '../../components/Layout/Layout'
import { getSalaryFromApiByUuid } from '../../consumers/api'
import { SalaryReportPage } from '../../components/Pages/SalaryReportPage'

export default function SalaryReport() {
  const router = useRouter()
  const { uuid } = router.query

  const [salaryExchange, setSalaryExchange] = useState({})
  useEffect(() => {
    if (uuid) getSalaryFromApiByUuid(uuid, true).then(setSalaryExchange)
  }, [uuid])

  return (
    <div>
      <Layout title="Salary report">
        {salaryExchange?.salary?.id ? (
          <SalaryReportPage salaryExchange={salaryExchange} />
        ) : null}
      </Layout>
    </div>
  )
}
