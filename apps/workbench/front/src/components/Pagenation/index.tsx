import React from 'react'
//
import './index.less'
//
import { ReactComponent as Left } from '@assets/comps/pagen-left.svg'
import { ReactComponent as Right } from '@assets/comps/pagen-right.svg'
import { IPaginationProps } from '@/interface/some'



const Pagination: React.FC<IPaginationProps> = (props) => {
  const pageTemp = Array.from({ length: props.total }, (_, i) => i + 1)
  const pagenumTemp = pageTemp.filter(i => i <= props.pages)

  // const visPage = props.pages
  const [pages, setPages] = React.useState<number[]>([])
  const [total, setTotal] = React.useState<number>(props.total)
  // const [visPage, setVisPage] = React.useState(props.pages)
  const [current, setCurrent] = React.useState(props.current)
  const [isFarStart, setIsFarStart] = React.useState(false)
  const [isFarEnd, setIsFarEnd] = React.useState(false)
  const [pagenum, setPagenum] = React.useState<number[]>([])

  const step = () => {
    return Math.floor(props.pages / 2)
  }

  React.useEffect(() => {
    setTotal(props.total)
    // setVisPage(props.pages)
    // setCurrent(props.current)
    setPages(Array.from({ length: props.total }, (_, i) => i + 1))
    setPagenum(pagenumTemp)

  }, [props.pages, props.total])

  React.useEffect(() => {
    // setTotal(props.total)
    // setVisPage(props.pages)
    // setCurrent(props.current)
  }, [props.total])


  React.useEffect(() => {
    let temp: typeof pages = []
    if (current < props.pages / 2 + 1) {
      setIsFarStart(false)
      setIsFarEnd(true)
      temp = pages.filter(i => i <= props.pages)
    }
    if (current >= total - props.pages / 2 - 1) {
      setIsFarStart(true)
      setIsFarEnd(false)
      temp = pages.filter(i => i > total - props.pages)
    }
    if (current >= props.pages / 2 + 1 && current < total - props.pages / 2 - 1) {
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
          {total > 5 && isFarStart && <div className="__pagination_swrapper_step __pagination_q">...</div>}
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
          {total > 5 && isFarEnd && <div className="__pagination_swrapper_step __pagination_q">...</div>}
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