import { render } from '@testing-library/react'
import App from '../src/App.jsx'
import PrivacyP from '../src/layout/privacyPolicy.jsx'

//router

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