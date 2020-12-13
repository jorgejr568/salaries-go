import { useState, useEffect } from 'react'

export function YourSalariesPage() {
  const [salaries, setSalaries] = useState([])

  useEffect(() => {
    const lsSalaries = JSON.parse(localStorage.getItem('mySalaries'))
    if (Array.isArray(lsSalaries)) setSalaries(lsSalaries)
  }, [])

  return (
    <div className="container">
      <h2>Your salaries</h2>
      <hr />
      <div className="row">
        {salaries.map((salary, key) => (
          <SalaryCard salary={salary} key={key} />
        ))}
      </div>
    </div>
  )
}

import { SalaryCard } from '../Salary/SalaryCard'
