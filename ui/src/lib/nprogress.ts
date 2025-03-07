import { debounce } from 'lodash-es'
import NProgress from 'nprogress'

NProgress.configure({ showSpinner: false, trickleSpeed: 300 })

const done = debounce(NProgress.done, 300, {
  leading: false,
  trailing: true,
})

export function useNProgress() {
  return {
    start: NProgress.start,
    done,
  }
}
