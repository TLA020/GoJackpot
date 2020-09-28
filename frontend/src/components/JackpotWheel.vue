<template>
  <div id="wrapper">
    <canvas ref="canvas"></canvas>
    <img src="../assets/pointer.png" id="pointer" alt="Pointer" :style="tickerRotationDegrees" />
  </div>
</template>

<script>
import {
  degreesToRadians,
  radiansToDegrees,
  lengthDirX,
  lengthDirY,
  percentageToRadians,
  resetEase,
  spinEase
} from "../helpers/JackpotHelpers";
import jackpotMixin from "../mixins/jackpot.mixin";

export default {
  name: "JackpotWheel",
  mixins: [jackpotMixin],

  data: () => ({
    index: 0,
    ticker: require("../assets/pointer.png")
  }),

  mounted() {
    this.canvas = this.$refs.canvas.getContext("2d");

    const { width, height } = this.$el.getClientRects()[0];

    this.width = width;
    this.height = height;

    this.sliceMiddle = width / 2 - this.sliceWidth / 2;

    this.$refs.canvas.width = width;
    this.$refs.canvas.height = height;

    const ticker = new Image();

    ticker.onload = () => {
      ticker.width = 48;
      this.ticker = ticker;
      this.draw();
    };

    ticker.src = this.ticker;
  },

  methods: {
    spin(percentage) {
      // const { startRadians, difference } = this.getSliceRangeOnIndex(
      //   this.index
      // );

      const targetRadian = -percentageToRadians(percentage) + Math.PI * 2 * 10;
      // const targetRadian = -startRadians - difference / 2 + Math.PI * 2 * 10;

      this.radiansExtra = this.radiansExtra % (Math.PI * 2);

      /*setTimeout(() => {
        this.radiansExtra = -this.radiansExtra % Math.PI * 2;
      }, 1000);*/

      this.state = "resetting";
      this.animate(
        Date.now(),
        true,
        resetEase,
        1000,
        this.radiansExtra,
        degreesToRadians(-120)
      ).then(() => {
        setTimeout(() => {
          this.state = "spinning";
          this.animate(Date.now(), false, spinEase, 10000, 0, targetRadian).then(() => {
            this.$emit("stopped");
            this.state = "winner-picked";
          });
        }, 1000);
      });
    },

    reset() {
      this.state = "idle";
    },

    draw() {
      this.canvas.clearRect(0, 0, this.width, this.height);

      if (this.tickerRotation > 0) {
        const difference = (1 / 350) * Math.min(350, (Date.now() - this.tickerHit));

        this.tickerRotation = Math.max(this.tickerRotation - resetEase(difference), 0);
      }

      if (this.backgroundColor) {
        this.canvas.fillStyle = this.backgroundColor;
        this.canvas.rect(0, 0, this.width, this.height);
        this.canvas.fill();
      }

      let radiansOffset = +(this.radiansOffset + this.radiansExtra).toFixed(4);

      this.pool.forEach((bet, index) => {
        const timeDifference = Date.now() - bet.addedAt;
        const percentage = Math.min(1, (1 / 500) * timeDifference);
        const radians = +(
          this.getSliceRadiansOnIndex(index) * percentage
        ).toFixed(4);
        const middleRadian = radiansOffset + radians / 2;

        this.canvas.beginPath();
        this.canvas.strokeStyle = this.getColor(index);
        this.canvas.arc(
          this.width / 2,
          this.height / 2,
          (this.width / 2) * 0.8,
          radiansOffset,
          radiansOffset + radians
        );
        this.canvas.lineWidth = this.sliceWidth;
        this.canvas.stroke();

        const difference = radiansOffset % (Math.PI * 2);

        if (this.radiansOffset - difference > 0 && this.radiansOffset - difference < 0.05) {
          this.tickerHit = Date.now();
          this.tickerRotation = Math.min(60, this.tickerRotation + 10);
        }

        if (radians > 0.26 && bet.image) {
          this.canvas.drawImage(
            bet.image,
            this.width / 2 +
              lengthDirX(this.sliceMiddle, middleRadian) -
              bet.image.width / 4,
            this.height / 2 +
              lengthDirY(this.sliceMiddle, middleRadian) -
              bet.image.height / 4,
            bet.image.width / 2,
            bet.image.height / 2
          );

          /*this.canvas.fillText(
            difference.toFixed(2).toString() + " " + this.radiansOffset.toFixed(2).toString(),
            this.width / 2 +
            lengthDirX(this.sliceMiddle, middleRadian) -
            bet.image.width / 4,
            this.height / 2 +
            lengthDirY(this.sliceMiddle, middleRadian) -
            bet.image.height / 4,);*/
        }

        radiansOffset += radians;
      });

      this.canvas.textAlign = "center";
      this.canvas.textBaseline = "middle";

      if (typeof this.drawTotal === "function") {
        this.drawTotal(
          this.canvas,
          this.width / 2,
          this.height / 2,
          this.total
        );
      } else {
        this.canvas.fillStyle = this.fontColor ?? "black";
        this.canvas.font = this.font ?? "bold 24px serif";
        this.canvas.fillText(
          `$${this.total.toFixed(2)}`,
          this.width / 2,
          this.height / 2
        );

        if (this.timeLeft && this.state !== "winner-picked") {
          this.canvas.font = this.font ?? "bold 36px serif";
          this.canvas.fillText(`${this.timeLeft <= 1 ? "GET READY!" : this.timeLeft}`, this.width / 2, this.height / 2 - 40)
        }

        if (this.state === "winner-picked" && this.winnerText) {
          this.canvas.font = this.font ?? "25px serif";
          this.canvas.fillText(this.winnerText, this.width / 2, this.height / 2 - 40)
        }
      }

      // this.canvas.drawImage(this.ticker, (this.width / 2) - (this.ticker.width / 2), 10);

      if (this.idleSpeed > 0 && this.state === "idle") {
        this.radiansExtra += degreesToRadians(this.idleSpeed);
      }

      requestAnimationFrame(this.draw);
    },

    getSliceRadiansOnIndex(index) {
      const share = ((100 / this.total) * this.pool[index].amount) / 100;
      const degrees = 360 * share;

      return degreesToRadians(degrees);
    },

    getSliceRangeOnIndex(index) {
      let startRadians = 0;

      for (let i = 0; i < index; i++) {
        startRadians += this.getSliceRadiansOnIndex(i);
      }

      let endRadians = startRadians + this.getSliceRadiansOnIndex(index);

      return {
        startRadians,
        endRadians,
        difference: Math.abs(endRadians - startRadians)
      };
    },

    animate(
      startTime,
      isNegative,
      ease,
      duration,
      startRadian,
      targetRadian,
      resolve
    ) {
      let onDone;

      if (!resolve) {
        onDone = new Promise(r => (resolve = r));
      }

      const now = Date.now();
      let t = duration > 0 ? (now - startTime) / duration : 1;
      const difference = Math.abs(startRadian - targetRadian);

      if (isNegative) {
        this.radiansExtra = startRadian - difference * ease(t);
      } else {
        this.radiansExtra = startRadian + difference * ease(t);
      }

      if (t >= 1) {
        this.radiansExtra = targetRadian;
        resolve();
      } else {
        requestAnimationFrame(() =>
          this.animate(
            startTime,
            isNegative,
            ease,
            duration,
            startRadian,
            targetRadian,
            resolve
          )
        );
      }

      return onDone;
    }
  },

  computed: {
    radiansToDegrees() {
      return radiansToDegrees;
    },

    tickerRotationDegrees() {
      return {
        transform: `rotate(${-this.tickerRotation}deg)`
      };
    },

    degreesToRadians() {
      return degreesToRadians;
    }
  }
};
</script>

<style scoped>
#wrapper {
  position: relative;
  width: 100%;
  height: 100%;
}

#pointer {
  position: absolute;
  left: calc(50% - 24px);
  top: -35px;
}
</style>
