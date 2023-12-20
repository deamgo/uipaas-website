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
  const EL = render(<RouterProvider router={router} />)
  expect(EL).toMatchSnapshot()
})

it('Sign In page Test', () => {
  const router = createMemoryRouter(
    routes,
    {
      initialEntries: ['/s'],
      initialIndex: 1,
    }
  )
  const EL = render(<RouterProvider router={router} />)
  expect(EL).toMatchSnapshot()
})

it('Sign Emial verifi page Test', () => {
  const router = createMemoryRouter(
    routes,
    {
      initialEntries: ['/s/ev'],
      initialIndex: 1,
    }
  )
  const EL = render(<RouterProvider router={router} />)
  expect(EL).toMatchSnapshot()
})

it('Sign Reset pwd page Test', () => {
  const router = createMemoryRouter(
    routes,
    {
      initialEntries: ['/s/ryp'],
      initialIndex: 1,
    }
  )
  const EL = render(<RouterProvider router={router} />)
  expect(EL).toMatchSnapshot()
})


it('Sign Reset pwd page Test', () => {
  const router = createMemoryRouter(
    routes,
    {
      initialEntries: ['/privacy'],
      initialIndex: 1,
    }
  )
  const EL = render(<RouterProvider router={router} />)
  expect(EL).toMatchSnapshot()
})

it('Apps page Test', () => {
  const router = createMemoryRouter(
    routes,
    {
      initialEntries: ['/'],
      initialIndex: 1,
    }
  )
  const EL = render(<RouterProvider router={router} />)
  expect(EL).toMatchSnapshot()
})

it('User page Test', () => {
  const router = createMemoryRouter(
    routes,
    {
      initialEntries: ['/u'],
      initialIndex: 1,
    }
  )
  const EL = render(<RouterProvider router={router} />)
  expect(EL).toMatchSnapshot()
})