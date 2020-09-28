import {
  createRoundedImage,
  degreesToRadians
} from "../helpers/JackpotHelpers";

export default {
  props: {
    pool: {
      type: Array,
      required: true,
      default: () => []
    },

    timeLeft: {
      type: Number,
      default: () => 0
    },

    winnerText: {
      type: String,
      default: () => null
    },

    offsetDegrees: {
      type: Number,
      default: () => 0
    },

    idleSpeed: {
      type: Number,
      default: () => 0
    },

    bgColor: String,

    font: String,

    fontColor: String,

    chartWidth: Number,

    drawTotal: Object
  },

  data: () => ({
    width: 0,
    height: 0,
    state: "idle",
    sliceMiddle: 0,
    canvas: null,
    radiansOffset: 0,
    radiansExtra: 0,
    tickerRotation: 0,
    tickerHit: 0,
    colors: [
      "#620F17",
      "#5C2056",
      "#2E102C",
      "#F7B000",
      "#E76E35",
      "#CB4856",
      "#BF1E2E",
      "#961823"
    ]
  }),

  mounted() {
    this.radiansOffset = degreesToRadians(this.offsetDegrees);
  },

  methods: {
    getColor(index) {
      const targetIndex = index % this.colors.length;

      if (targetIndex >= this.colors.length) {
        return this.getColor(targetIndex);
      }

      return this.colors[targetIndex];
    }
  },

  computed: {
    total() {
      return this.pool.reduce((a, b) => a + b.amount, 0);
    },

    backgroundColor() {
      return this.bgColor;
    },

    sliceWidth() {
      return this.chartWidth ?? 60;
    }
  },

  watch: {
    pool: {
      immediate: true,
      handler(pool) {
        pool.forEach(bet => {
          if (!bet.addedAt) {
            bet.addedAt = Date.now();
          }

          if (!bet.image) {
            createRoundedImage(bet.avatar).then(image => {
              bet.image = image;
            });
          }
        });
      }
    }
  }
};
