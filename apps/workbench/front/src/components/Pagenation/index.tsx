import React from 'react'
//
import './index.less'
//
import { ReactComponent as Left } from '@assets/comps/pagen-left.svg'
import { ReactComponent as Right } from '@assets/comps/pagen-right.svg'
import { IPaginationProps } from '@/interface/some'



const Pagination: React.FC<IPaginationProps> = (props) => {

  // const visPage = props.pages
  const [pages, setPages] = React.useState(Array.from({ length: props.total }, (_, i) => i + 1))
  const [total, setTotal] = React.useState(props.total)
  const [visPage, setVisPage] = React.useState(props.pages)
  const [current, setCurrent] = React.useState(props.current)
  const [isFarStart, setIsFarStart] = React.useState(false)
  const [isFarEnd, setIsFarEnd] = React.useState(false)
  const [pagenum, setPagenum] = React.useState(pages)

  const step = () => {
    return Math.floor(visPage / 2)
  }

  React.useEffect(() => {
    setPages(Array.from({ length: props.total }, (_, i) => i + 1))
    setTotal(props.total)
    setVisPage(props.pages)
    setCurrent(props.current)
  }, [props])


  React.useEffect(() => {
    let temp: typeof pages = []
    if (current < visPage / 2 + 1) {
      setIsFarStart(false)
      setIsFarEnd(true)
      temp = pages.filter(i => i <= visPage)
    }
    if (current >= total - visPage / 2 - 1) {
      setIsFarStart(true)
      setIsFarEnd(false)

      temp = pages.filter(i => i > total - visPage)
    }
    if (current >= visPage / 2 + 1 && current < total - visPage / 2 - 1) {
      setIsFarEnd(true)
      setIsFarStart(true)
      temp = pages.filter(i => i > current - step() && i <= current + step())

    }

    setPagenum(temp)

    props.onCurrentPageChange(current)
    return () => { }
  }, [current])



  return (
    <>
      <div className="__pagination">
        <div className="__pagination_swrapper">
          <div className="__pagination_swrapper_total">
            <span>Total&nbsp;{props.total}</span>
          </div>
          <button
            disabled={current < 2}
            className="__pagination_swrapper_row __pagination_q __pagination_disable"
            onClick={() => setCurrent(c => c - 1)}>
            <Left />
          </button>
          {isFarStart && <div className="__pagination_swrapper_step __pagination_q">...</div>}
          {
            pagenum.map((item, index) => (
              <div
                key={index}
                className={`__pagination_swrapper_step __pagination_q ${current === item && '__pagination_current'}`}
                onClick={() => setCurrent(item)} >
                {item}
              </div>
            ))
          }
          {isFarEnd && <div className="__pagination_swrapper_step __pagination_q">...</div>}
          <button
            disabled={current >= total}
            className="__pagination_swrapper_row __pagination_q __pagination_disable"
            onClick={() => setCurrent(c => c + 1)}>
            <Right />
          </button>
        </div>
      </div>
    </>
  )
}

export default Pagination