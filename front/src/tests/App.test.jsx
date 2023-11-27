import React from 'react'
import { render } from '@testing-library/react'
import App from '../App.jsx'
import Err from '../layout/errorPage.jsx'
import PrivacyP from '../layout/Privacy-policy.jsx'
import { CollectInfo } from '../components/CollectInfo/index.jsx'

//router
import { createMemoryRouter, RouterProvider } from 'react-router-dom'


describe('Main', () => {
    it('App correct render', () => {
        const { baseEL } = render(<App />)
        expect(baseEL).toMatchSnapshot()
    })
    it('Privacy-Policy correct render', () => {
        const { baseEL } = render(<PrivacyP />)
        expect(baseEL).toMatchSnapshot()
    })
    it('Err page correct render', () => {
        const routes = [
            {
                path: '/',
                errorElement: <Err />   
            },
        ]
        const router = createMemoryRouter(
            routes,
            {
                initialEntries: ['/eee'],
                initialIndex: 1,
            }
        )
        const { baseEL } = render(<RouterProvider router={router}><App /></RouterProvider>)
        expect(baseEL).toMatchSnapshot()
    })


    test('calls onClick function when clicked', () => {
        const mockOnClick = jest.fn();
        const { asFragment } = render(<CollectInfo/>);
        expect(asFragment()).toMatchSnapshot();
      });
})