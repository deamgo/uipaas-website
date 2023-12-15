import { RouterProvider, createMemoryRouter } from 'react-router-dom'

import { routes } from '@/router/router'
import { render } from '@testing-library/react'

it('Sign page Test', () => {
  const router = createMemoryRouter(
    routes,
    {
      initialEntries: ['/s/up'],
      initialIndex: 1,
    }
  )
  const SignIN = render(<RouterProvider router={router} />)
  expect(SignIN).toMatchSnapshot()
})