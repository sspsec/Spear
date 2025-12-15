<template>
  <div class="watermark">
    <canvas ref="canvas" class="watermark-canvas"></canvas>
  </div>
</template>

<script>
export default {
  name: 'Watermark',
  props: {
    text: {
      type: String,
      default: ''
    },
    opacity: {
      type: Number,
      default: 0.1
    }
  },
  mounted() {
    this.drawWatermark()
    window.addEventListener('resize', this.drawWatermark)
  },
  beforeUnmount() {
    window.removeEventListener('resize', this.drawWatermark)
  },
  methods: {
    drawWatermark() {
      const canvas = this.$refs.canvas
      const ctx = canvas.getContext('2d')
      
      // 设置画布大小为窗口大小
      canvas.width = window.innerWidth
      canvas.height = window.innerHeight
      
      // 设置水印样式
      ctx.font = '16px var(--font-family)'
      ctx.fillStyle = `rgba(255, 255, 255, ${this.opacity})`
      ctx.rotate(-15 * Math.PI / 180) // 旋转 -15 度
      
      // 计算水印间距
      const text = this.text
      const textWidth = ctx.measureText(text).width
      const xGap = textWidth + 100
      const yGap = 100
      
      // 绘制水印
      for (let y = -canvas.height; y < canvas.height * 2; y += yGap) {
        for (let x = -canvas.width; x < canvas.width * 2; x += xGap) {
          ctx.fillText(text, x, y)
        }
      }
    }
  }
}
</script>

<style scoped>
.watermark {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  pointer-events: none;
  z-index: -1;
}

.watermark-canvas {
  width: 100%;
  height: 100%;
}
</style>