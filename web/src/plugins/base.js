import Vue from 'vue'
import BaseBtn from '@/components/base/Btn'
import BaseCard from '@/components/base/Card'
import BaseSubheading from '@/components/base/Subheading'

import BaseHeading from '@/components/base/Heading'
import BaseText from '@/components/base/Text'
import BaseBubble1 from '@/components/base/Bubble1'
import BaseBubble2 from '@/components/base/Bubble2'

Vue.component(BaseBtn.name, BaseBtn)
Vue.component(BaseCard.name, BaseCard)
Vue.component(BaseSubheading.name, BaseSubheading)

Vue.component('BaseHeading', BaseHeading)
Vue.component('BaseText', BaseText)
Vue.component('BaseBubble1', BaseBubble1)
Vue.component('BaseBubble2', BaseBubble2)
