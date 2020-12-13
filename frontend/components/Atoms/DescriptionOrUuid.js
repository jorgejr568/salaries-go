export function DescriptionOrUuidFn(salary) {
  return salary.description?.length ? salary.description : salary.uuid
}

export function DescriptionOrUuid({ salary }) {
  return <>{DescriptionOrUuidFn(salary)}</>
}
