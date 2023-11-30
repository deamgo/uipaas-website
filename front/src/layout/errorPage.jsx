import './err.less'

import { useRouteError, createBrowserRouter } from 'react-router-dom'

export default function ErrorPage() {
    // const error = useRouteError();
    const err = useRouteError()
  
    return (
      <div id="error-page">
        <h1>Oops!</h1>
        <p>Sorry, an unexpected error has occurred.</p>
        <p>
          <i>{err.statusText || err.message}</i>
        </p>
      </div>
    );
}
