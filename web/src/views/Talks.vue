<template>
  <div>
    <banner />
    <dates />
    <videos />
  </div>
</template>

<script>
export default {
  components: {
    banner: () => import('@/components/talks/Banner'),
    dates: () => import('@/components/talks/Dates'),
    videos: () => import('@/components/talks/Videos')
  }
}
</script>
