// import Vue from "vue";
//
// const socketEvents = {};
//
// Vue.prototype.$socketEvent = (event, fn) => {
//   if (!socketEvents[event]) {
//     socketEvents[event] = [];
//   }
//
//   socketEvents[event].push(fn);
// };
//
// export default store => {
//   store.subscribe(mutation => {
//     if (mutation.type === "SOCKET_ONMESSAGE") {
//       const { name: event, data } = mutation.payload;
//
//       if (socketEvents[event]) {
//         socketEvents[event].forEach(fn => fn(data));
//       }
//     }
//   });
// };
