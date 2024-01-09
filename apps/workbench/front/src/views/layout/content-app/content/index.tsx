import React from 'react';
//style
import './index.less'
//
import { ReactComponent as NoneContent } from '@assets/default/none-content.svg'
import { applicationStore } from '@/store/application';
import ApplicationBox from '@/components/ApplicationBox';
import { currentWorkspaceStore, wsStore } from '@/store/wsStore';
import { observer } from 'mobx-react-lite';

interface IApplication {
  created_by: number
  deleted_by: number
  description: string
  icon: string
  id: string
  name: string
  status: number
  workspace_id: string
}

// interface IContentAppContentProps {
//   shouldRefresh: number
// }

const Content: React.FC = () => {
  const [isEmpty, setIsEmpty] = React.useState(false)
  const [appList, setAppList] = React.useState<IApplication[]>([])

  React.useEffect(() => {
    if (!currentWorkspaceStore.getCurrentWorkspace()) {
      currentWorkspaceStore.setCurrentWorkspace(wsStore.getWsList()[0])
    }
    (async function () {
      await getAppList()
    }())
  }, [applicationStore.getAppStats().shouldRefresh])

  React.useEffect(() => {
    if (appList && appList.length > 0) {
      setIsEmpty(false)
    } else {
      setIsEmpty(true)
    }
  }, [appList])

  const getAppList = async () => {
    const data = await applicationStore.getApp() as IApplication[]
    setAppList(data)
  }

  // const memorizeGetAppList = React.useMemo(() => getAppList, [applicationStore.getAppStats().shouldRefresh])

  return (
    <>
      <div className="__appcontent">
        {isEmpty ? (
          <>
            <div className="__appcontent_empty">
              <div className="__appcontent_empty_svg">
                <NoneContent />
              </div>
              <span className="__appcontent_empty_span">
                No content, please create
              </span>
            </div>
          </>
        ) : (
          <>
            <div className="__appcontent_apps">
              {appList && appList.map(item => (
                <>
                  <ApplicationBox id={item.id} key={item.id} name={item.name} desc={item.description} label={item.status} />
                </>
              ))}
            </div>
          </>
        )}
      </div>
    </>
  )
}

export default observer(Content)