<template>
  <div>
    <uploader />
  </div>
</template>

<script>
export default {
  components: {
    uploader: () => import('@/components/upload/Uploader'),
  }
}
</script>
