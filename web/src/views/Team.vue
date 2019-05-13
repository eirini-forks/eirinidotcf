<template>
  <div>
    <banner />
    <team />
  </div>
</template>

<script>
export default {
  components: {
    team: () => import('@/components/team/Avatars'),
    banner: () => import('@/components/team/Banner'),
  }
}
</script>
