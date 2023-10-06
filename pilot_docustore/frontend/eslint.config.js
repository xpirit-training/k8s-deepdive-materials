import antfu from '@antfu/eslint-config'
import unocss from '@unocss/eslint-plugin'

export default antfu(
  {
    root: true,
  },
  unocss.configs.flat,
  
)
