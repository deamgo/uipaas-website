import { render } from '@testing-library/react'
import App from '../App.jsx'
import PrivacyP from '../layout/privacyPolicy.jsx'

//router
import { createMemoryRouter, RouterProvider, Navigate } from 'react-router-dom'


describe('Main', () => {
    it('App correct render', () => {
        const { baseEL } = render(<App />)
        expect(baseEL).toMatchSnapshot()
    })
    it('Privacy-Policy correct render', () => {
        const { baseEL } = render(<PrivacyP />)
        expect(baseEL).toMatchSnapshot()
    })
})