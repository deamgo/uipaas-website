import { RouterProvider, createMemoryRouter } from 'react-router-dom'

import { routes } from '@/router/router'
import { render } from '@testing-library/react'

it('Sign Up page Test', () => {
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

it('Sign In page Test', () => {
  const router = createMemoryRouter(
    routes,
    {
      initialEntries: ['/s/in'],
      initialIndex: 1,
    }
  )
  const SignIN = render(<RouterProvider router={router} />)
  expect(SignIN).toMatchSnapshot()
})

it('Sign Emial verifi page Test', () => {
  const router = createMemoryRouter(
    routes,
    {
      initialEntries: ['/s/ev'],
      initialIndex: 1,
    }
  )
  const SignIN = render(<RouterProvider router={router} />)
  expect(SignIN).toMatchSnapshot()
})

it('Sign Reset pwd page Test', () => {
  const router = createMemoryRouter(
    routes,
    {
      initialEntries: ['/s/ryp'],
      initialIndex: 1,
    }
  )
  const SignIN = render(<RouterProvider router={router} />)
  expect(SignIN).toMatchSnapshot()
})