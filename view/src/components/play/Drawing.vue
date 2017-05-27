<template lang="pug">
  canvas#drawing(@mousedown="mouseDown",
    @mouseup="mouseUp",
    @mousemove="mouseMove",
    @contextmenu="contextMenu")
</template>

<script>
  import Vue from 'vue';
  import debounce from 'debounce';

  export default {
    name: 'ss-drawing',

    data() {
      return {
        isPressed: false,
        canvasWidth: 0,
        canvasHeight: 0,
        boundingLeft: 0,
        boundingTop: 0,
        thickness: 5,
        lines: [],
        tempLine: [],
        ctx: null,
      };
    },

    computed: {
      colour() {
        return this.$store.getters.colour;
      },
      width() {
        return this.canvasWidth;
      },
      height() {
        return this.canvasHeight;
      },
    },

    mounted() {
      this.$nextTick(() => {
        // Add resize event handling to canvas
        window.addEventListener('resize', this.onResize);
        window.addEventListener('scroll', this.onScroll);
        // Initialize canvas
        this.ctx = this.$el.getContext('2d');
        this.$el.addEventListener('touchstart', this.touchDown);
        this.$el.addEventListener('touchmove', this.touchMove);
        this.$el.addEventListener('touchend', this.mouseUp);
        this.$el.addEventListener('touchcancel', this.mouseUp);
        this.onResize();
      });
    },

    methods: {
      // Resize with debounce delay of 100ms to prevent the event from firing too often
      onResize: debounce(function resize() {
        // On resize, update canvas properties and redraw all lines
        this.updateCanvasProps();
        this.redraw();
      }, 100),

      onScroll: debounce(function scroll() {
         // On scroll, update canvas bounds
        this.updateCanvasBounds();
      }, 100),

      updateCanvasBounds() {
        // Update x, y anchor coordinates
        const boundingRect = this.$el.getBoundingClientRect();
        this.boundingTop = boundingRect.top;
        this.boundingLeft = boundingRect.left;
      },

      updateCanvasProps() {
        // Update canvas size
        this.canvasWidth = this.$el.clientWidth;
        this.canvasHeight = this.$el.clientHeight;
        Vue.set(this.$el, 'width', this.canvasWidth);
        Vue.set(this.$el, 'height', this.canvasHeight);

        this.updateCanvasBounds();

        // Update canvas context
        this.thickness = 5 * (this.canvasWidth / 1000);
        this.ctx.lineWidth = this.thickness;
      },

      draw() {
        const ctx = this.ctx;
        ctx.strokeStyle = this.colour;
        ctx.lineJoin = 'round';
        ctx.lineCap = 'round';

        for (let i = 0; i < this.tempLine.length; i += 1) {
          ctx.beginPath();
          if (this.tempLine[i] && i) {
            ctx.moveTo(this.tempLine[i - 1].x, this.tempLine[i - 1].y);
          } else {
            ctx.moveTo(this.tempLine[i].x - 1, this.tempLine[i].y);
          }
          ctx.lineTo(this.tempLine[i].x, this.tempLine[i].y);
          ctx.closePath();
          ctx.stroke();
        }
      },

      redraw() {
        const ctx = this.ctx;
        ctx.lineJoin = 'round';
        ctx.lineCap = 'round';

        for (let n = 0; n < this.lines.length; n += 1) {
          const lineData = this.lines[n];
          const thickness = lineData.thickness;
          const line = lineData.line;
          const scaleX = this.canvasWidth / lineData.sizeX;
          const scaleY = this.canvasHeight / lineData.sizeY;
          ctx.lineWidth = thickness * scaleX;

          ctx.strokeStyle = lineData.colour;
          if (line && line[0]) {
            ctx.beginPath();
            ctx.moveTo(line[0].x * scaleX, line[0].y * scaleY);

            // Use quadratic bezier interpolation for smoother lines
            if (line[1]) {
              let i = 1;
              for (; i < line.length - 2; i += 1) {
                const xc = (line[i].x + line[i + 1].x) / 2;
                const yc = (line[i].y + line[i + 1].y) / 2;
                ctx.quadraticCurveTo(line[i].x * scaleX, line[i].y * scaleY,
                  xc * scaleX, yc * scaleY);
              }
              // Curve through the last two points
              ctx.quadraticCurveTo(line[i].x * scaleX, line[i].y * scaleY,
                line[i + 1].x * scaleX, line[i + 1].y * scaleY);
              ctx.stroke();
              ctx.closePath();
            } else {
              ctx.lineTo((line[0].x * scaleX) - 1, line[0].y * scaleY);
              ctx.closePath();
              ctx.stroke();
            }
          }
        }
      },

      touchMove(event) {
        event.preventDefault();
        if (this.colour) {
          // Touch down
          this.isPressed = true;

          // Add initial (x, y) point
          this.tempLine.push({
            x: event.touches[0].clientX - this.boundingLeft,
            y: event.touches[0].clientY - this.boundingTop,
          });

          // Draw
          this.draw();
        }
      },

      touchDown(event) {
        event.preventDefault();
        if (this.isPressed) {
          // Add (x, y) point
          this.tempLine.push({
            x: event.touches[0].clientX - this.boundingLeft,
            y: event.touches[0].clientY - this.boundingTop,
          });
          this.draw();
        }
      },

      mouseMove(event) {
        event.preventDefault();
        if (this.isPressed) {
          // Add (x, y) point
          this.tempLine.push({ x: event.clientX - this.boundingLeft, y: event.clientY - this.boundingTop });
          this.draw();
        }
      },

      mouseDown(event) {
        event.preventDefault();
        if (this.colour) {
          // Mouse is pressed
          this.isPressed = true;

          // Add initial (x, y) point
          this.tempLine.push({ x: event.clientX - this.boundingLeft, y: event.clientY - this.boundingTop });

          // Draw
          this.draw();
        }
      },

      mouseUp() {
        this.isPressed = false;
        if (this.tempLine.length > 0) {
          const lineData = {
            line: this.tempLine,
            colour: this.colour,
            sizeX: this.canvasWidth,
            sizeY: this.canvasHeight,
            thickness: this.thickness,
          };
          this.lines.push(lineData);
          this.tempLine = [];
        }
        this.ctx.clearRect(0, 0, this.canvasWidth, this.canvasHeight);
        this.redraw();
      },

      // Disable right-click context menu for canvas
      contextMenu(event) {
        event.preventDefault();
        return false;
      },
    },
  };
</script>

<style lang="stylus">
  #drawing
    width: 100%
    height: 100%
    cursor: default
    user-select: none

  .debug
    position: absolute
    top: 0
    left: 0
</style>
