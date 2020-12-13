import Link from 'next/link'
import { useRouter } from 'next/router'

export function Navbar() {
  const router = useRouter()
  const activeClass = (href) => {
    return router.pathname === href ? 'active' : ''
  }
  return (
    <nav className="navbar navbar-expand-lg navbar-dark bg-dark">
      <div className="container">
        <Link href="/">
          <a className="navbar-brand" href="#">
            SalariesAPI
          </a>
        </Link>
        <button
          className="navbar-toggler"
          type="button"
          data-toggle="collapse"
          data-target="#navbarNavAltMarkup"
          aria-controls="navbarNavAltMarkup"
          aria-expanded="false"
          aria-label="Toggle navigation"
        >
          <span className="navbar-toggler-icon" />
        </button>
        <div className="collapse navbar-collapse" id="navbarNavAltMarkup">
          <div className="navbar-nav">
            <Link href="/">
              <a className={'nav-item nav-link ' + activeClass('/')} href="#">
                Your salaries
              </a>
            </Link>
            <Link href="/salary/attach">
              <a
                className={'nav-item nav-link ' + activeClass('/salary/attach')}
                href="#"
              >
                Attach salary
              </a>
            </Link>
            <Link href="/salary">
              <a
                className={'nav-item nav-link ' + activeClass('/salary')}
                href="#"
              >
                New salary
              </a>
            </Link>
          </div>
        </div>
      </div>
    </nav>
  )
}
