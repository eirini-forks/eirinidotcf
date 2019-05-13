<template>
  <div>
    <welcome />
    <services />
    <connect />
    <reads />
  </div>
</template>

<script>
export default {
  components: {
    Welcome: () => import('@/components/info/Welcome'),
    Connect: () => import('@/components/info/Connect'),
    Services: () => import('@/components/info/Services'),
    Reads: () => import('@/components/info/Reads')
  }
}
</script>
