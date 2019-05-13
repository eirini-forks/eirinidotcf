<template>
  <div id="Pheed">
    <articles>
      <banner />
    </articles>
  </div>
</template>

<script>
  export default {
    name: 'Pheed',

    components: {
      Articles: () => import('@/components/pheed/Articles'),
      Banner: () => import('@/components/pheed/Banner')
    }
  }
</script>
