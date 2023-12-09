<template>
  <header>
    <img class="utc" alt="utc logo" src="@/assets/utc_logo.jpg" />
    <h1 class="title">Responsive AI Clusters in Supply Chain</h1>
    <button @click="start" class="btn">Start</button>
  </header>
  <div class="information">
    <h2 class="date">Date: {{ date }}</h2>
    <!-- <h2 class="event">Event: {{ event }}</h2> -->
  </div>
  <div ref="visualization">
    <div class="table">
      <div class="warehousetable">
        <table>
          <thead>
            <tr>
              <th colspan="2">Warehouse</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="(value, key) in stock0" :key="key">
              <td>{{ value[0] }}</td>
              <td>{{ value[1] }}</td>
            </tr>
          </tbody>
        </table>
      </div>
      <div class="supermarkettable">
        <table>
          <thead>
            <tr>
              <th colspan="2">Auchan</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="(value, key) in stock1" :key="key">
              <td>{{ value[0] }}</td>
              <td>{{ value[1] }}</td>
            </tr>
          </tbody>
        </table>
      </div>
      <div class="supermarkettable">
        <table>
          <thead>
            <tr>
              <th colspan="2">Carrefour</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="(value, key) in stock2" :key="key">
              <td>{{ value[0] }}</td>
              <td>{{ value[1] }}</td>
            </tr>
          </tbody>
        </table>
      </div>
      <div class="supermarkettable">
        <table>
          <thead>
            <tr>
              <th colspan="2">Monoprix</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="(value, key) in stock3" :key="key">
              <td>{{ value[0] }}</td>
              <td>{{ value[1] }}</td>
            </tr>
          </tbody>
        </table>
      </div>
      <div class="supermarkettable">
        <table>
          <thead>
            <tr>
              <th colspan="2">Normal</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="(value, key) in stock4" :key="key">
              <td>{{ value[0] }}</td>
              <td>{{ value[1] }}</td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
    <div class="communication">
      <button @click="toggleMessages1(1)" class="toggle-button show1">
        Auchan Messages
      </button>
      <button @click="toggleMessages2(1)" class="toggle-button show2">
        Carrefour Messages
      </button>
      <button @click="toggleMessages3(1)" class="toggle-button show3">
        Monoprix Messages
      </button>
      <button @click="toggleMessages4(1)" class="toggle-button show4">
        Normal Messages
      </button>
      <!-- <button @click="centralMessages1(1)" class="central-button centralshow1">Central to Auchan</button>
      <button @click="centralMessages2(1)" class="central-button centralshow2">Central to Carrefour</button>
      <button @click="centralMessages3(1)" class="central-button centralshow3">Central to Monoprix</button>
      <button @click="centralMessages4(1)" class="central-button centralshow4">Central to Normal</button> -->
      <!-- supermarkets messages -->
      <!-- (message1 || message5) &&  -->
      <div class="event">
        <p class="event1">{{ event1 }}</p>
        <p class="event2">{{ event2 }}</p>
        <p class="event3">{{ event3 }}</p>
        <p class="event4">{{ event4 }}</p>
      </div>

      <div
        ref="messages1"
        class="message-container1"
        v-if="(message1 || message5) && showMessages1"
      >
        <span
          v-for="msg in messages1_text"
          :key="msg.id"
          :class="['message', getMessageClass(msg.speakerid)]"
          >{{ msg.text }}
        </span>
      </div>
      <div
        ref="messages2"
        class="message-container2"
        v-if="(message2 || message6) && showMessages2"
      >
        <span
          v-for="msg in messages2_text"
          :key="msg.id"
          :class="['message', getMessageClass(msg.speakerid)]"
          >{{ msg.text }}
        </span>
      </div>
      <div
        ref="messages3"
        class="message-container3"
        v-if="(message3 || message7) && showMessages3"
      >
        <span
          v-for="msg in messages3_text"
          :key="msg.id"
          :class="['message', getMessageClass(msg.speakerid)]"
          >{{ msg.text }}
        </span>
      </div>
      <div
        ref="messages4"
        class="message-container4"
        v-if="(message4 || message8) && showMessages4"
      >
        <span
          v-for="msg in messages4_text"
          :key="msg.id"
          :class="['message', getMessageClass(msg.speakerid)]"
          >{{ msg.text }}
        </span>
      </div>
      <!-- warehouse messages -->
      <!-- <div ref="messages5" class="message-container5" v-if="message5 && showMessages5">
                <span v-for="msg in messages5_text" :key="msg.id">{{ msg.text }} </span>
            </div>
            <div ref="messages6" class="message-container6" v-if="message6 && showMessages6">
                <span v-for="msg in messages6_text" :key="msg.id">{{ msg.text }} </span>
            </div>
            <div ref="messages7" class="message-container7" v-if="message7 && showMessages7">
                <span v-for="msg in messages7_text" :key="msg.id">{{ msg.text }} </span>
            </div>
            <div ref="messages8" class="message-container8" v-if="message8 && showMessages8">
                <span v-for="msg in messages8_text" :key="msg.id">{{ msg.text }} </span>
            </div> -->
    </div>
  </div>
</template>

<script>
import * as d3 from "d3";

export default {
  name: "App",
  data() {
    return {
      stock0: new Map([
        ["Olive Oil", 0],
        ["Baguette", 0],
        ["Manchego Cheese", 0],
        ["Black Tea", 0],
      ]),
      stock1: new Map([
        ["Olive Oil", 0],
        ["Baguette", 0],
        ["Manchego Cheese", 0],
        ["Black Tea", 0],
      ]),
      stock2: new Map([
        ["Olive Oil", 0],
        ["Baguette", 0],
        ["Manchego Cheese", 0],
        ["Black Tea", 0],
      ]),
      stock3: new Map([
        ["Olive Oil", 0],
        ["Baguette", 0],
        ["Manchego Cheese", 0],
        ["Black Tea", 0],
      ]),
      stock4: new Map([
        ["Olive Oil", 0],
        ["Baguette", 0],
        ["Manchego Cheese", 0],
        ["Black Tea", 0],
      ]),
      onedaytime: 60000, // 60s/day
      date: null,
      event1: "No event",
      event2: "No event",
      event3: "No event No event",
      event4: "No event",
      outlet1: null,
      supermarketInfo1: null,
      outlet2: null,
      outlet3: null,
      outlet4: null,
      message1: null,
      message2: null,
      message3: null,
      message4: null,
      message5: null,
      message6: null,
      message7: null,
      message8: null,
      centralhub: null,
      svg: null,
      warehouse: { x: 1250, y: 500 },
      supermarkets: [
        { x: 550, y: 150 },
        { x: 550, y: 850 },
        { x: 1950, y: 150 },
        { x: 1950, y: 850 },
      ],
      messages1_text: [],
      messages2_text: [],
      messages3_text: [],
      messages4_text: [],
      // messages5_text: [],
      // messages6_text: [],
      // messages7_text: [],
      // messages8_text: [],
      nextMsgId1: 0,
      nextMsgId2: 0,
      nextMsgId3: 0,
      nextMsgId4: 0,
      // nextMsgId5: 0,
      // nextMsgId6: 0,
      // nextMsgId7: 0,
      // nextMsgId8: 0,
      showMessages1: true,
      showMessages2: true,
      showMessages3: true,
      showMessages4: true,
      // showMessages5: true,
      // showMessages6: true,
      // showMessages7: true,
      // showMessages8: true
    };
  },
  //////////
  mounted() {
    this.createVisualization();
  },
  //////////
  methods: {
    createVisualization() {
      const width = 2400;
      const height = 1200;
      // Create SVG canvas
      this.svg = d3
        .select(this.$refs.visualization)
        .append("svg")
        .attr("width", width)
        .attr("height", height);

      // 先定义一个道路图案
      this.svg
        .append("defs")
        .append("pattern")
        .attr("id", "roadPattern")
        .attr("width", 20) // 图案宽度，可以根据需要调整
        .attr("height", 20) // 图案高度，同上
        .attr("patternUnits", "userSpaceOnUse")
        .append("path")
        .attr("d", "M 10 0 L 10 20") // 创建线性图案
        .attr("stroke", "#666") // 道路颜色
        .attr("stroke-width", 6); // 道路宽度

      // Drawing the line between warehouse and supermarket
      this.supermarkets.forEach((s) => {
        this.svg
          // .append("line")
          // .style("stroke", "black")
          // .attr("x1", this.warehouse.x)
          // .attr("y1", this.warehouse.y)
          // .attr("x2", s.x)
          // .attr("y2", s.y + 15);
          .append("line")
          .style("stroke", "url(#roadPattern)") // 应用道路图案
          .style("stroke-width", 30) // 增加线条宽度
          .attr("x1", this.warehouse.x)
          .attr("y1", this.warehouse.y)
          .attr("x2", s.x)
          .attr("y2", s.y + 15)
          .attr("stroke-linecap", "round"); // 圆形线帽，使道路看起来更平滑
      });
      // warehouse
      this.svg
        .append("image")
        .attr("xlink:href", require("@/assets/warehouse.png"))
        .attr("x", this.warehouse.x - 175)
        .attr("y", this.warehouse.y - 195)
        .attr("width", 350)
        .attr("height", 350);
      // supermarket
      //const supermarketNames = ["Auchan", "Carrefour", "Monoprix", "Normal"];
      // this.supermarkets.forEach((s, index) => {
      //   this.svg
      //   .append("image")
      //   .attr("xlink:href", require("@/assets/supermarket.png"))
      //   .attr("x", s.x - 15)
      //   .attr("y", s.y - 15)
      //   .attr("width", 55)
      //   .attr("height", 55);
      //   //supermarket number
      //   this.svg
      //     .append("text")
      //     .attr("x", s.x - 30)
      //     .attr("y", s.y - 20)
      //     .text(`${supermarketNames[index]}`)
      //     .style("font-size", "28px")
      //     .style("font-family", "Arial, sans-serif")
      //     .attr("class", "supermarket-label");
      // });
      this.svg
        .append("image")
        .attr("xlink:href", require("@/assets/auchan.png"))
        .attr("x", this.supermarkets[0].x - 150)
        .attr("y", this.supermarkets[0].y - 150)
        .attr("width", 300)
        .attr("height", 300);
      this.svg
        .append("image")
        .attr("xlink:href", require("@/assets/carrefour.png"))
        .attr("x", this.supermarkets[1].x - 150)
        .attr("y", this.supermarkets[1].y - 150)
        .attr("width", 300)
        .attr("height", 300);
      this.svg
        .append("image")
        .attr("xlink:href", require("@/assets/monoprix.png"))
        .attr("x", this.supermarkets[2].x - 150)
        .attr("y", this.supermarkets[2].y - 150)
        .attr("width", 300)
        .attr("height", 300);
      this.svg
        .append("image")
        .attr("xlink:href", require("@/assets/normal.png"))
        .attr("x", this.supermarkets[3].x - 150)
        .attr("y", this.supermarkets[3].y - 150)
        .attr("width", 300)
        .attr("height", 300);
    },

    start() {
      this.$nextTick(() => {
        this.outlet1start();
        this.outlet2start();
        this.outlet3start();
        this.outlet4start();
        this.centralhubstart();
        this.message1start();
        this.message2start();
        this.message3start();
        this.message4start();
        this.message5start();
        this.message6start();
        this.message7start();
        this.message8start();
      });
    },

    ///////////
    centralhubstart() {
      this.centralhub = new WebSocket("ws://localhost:8080/centralhub");
      this.centralhub.onmessage = (event) => {
        const data = JSON.parse(event.data);
        this.freshInfo(data);
      };
    },

    freshInfo(generalInfo) {
      let date1 = generalInfo.date;
      let date0 = new Date(date1);
      var year = date0.getUTCFullYear();
      var month = date0.getUTCMonth() + 1; // 注意月份是从0开始的，因此要加1
      var day = date0.getUTCDate();
      this.date = this.pad(day) + "/" + this.pad(month) + "/" + year;
      ////////
      // this.event = generalInfo.event;
      Object.entries(generalInfo.event).forEach((table) => {
        let key = table[0];
        let value = table[1];
        switch (key) {
          case "1":
            this.event1 = value;
            break;
          case "2":
            this.event2 = value;
            break;
          case "3":
            this.event3 = value;
            break;
          case "4":
            this.event4 = value;
            break;
        }
      });
      this.stock0 = new Map(Object.entries(generalInfo.warehouseProduct));
    },

    pad(num) {
      return num < 10 ? "0" + num : num;
    },

    sendBoxToSupermarket(supermarketInfo) {
      //////////
      let target;
      const duration = this.onedaytime * supermarketInfo.deliveryTime;
      switch (supermarketInfo.id) {
        case "1":
          target = this.supermarkets[0];
          this.stock1 = new Map(Object.entries(supermarketInfo.productLeft));
          break;
        case "2":
          target = this.supermarkets[1];
          this.stock2 = new Map(Object.entries(supermarketInfo.productLeft));
          break;
        case "3":
          target = this.supermarkets[2];
          this.stock3 = new Map(Object.entries(supermarketInfo.productLeft));
          break;
        case "4":
          target = this.supermarkets[3];
          this.stock4 = new Map(Object.entries(supermarketInfo.productLeft));
          break;
      }
      //////////红蓝绿橙
      Object.entries(supermarketInfo.productAdd).forEach((table) => {
        let key = table[0];
        let value = table[1];
        let size = 5;
        if (value != 0) {
          if (key == "Olive Oil") {
            const boxa = this.svg
              .append("rect")
              .attr("x", this.warehouse.x + size)
              .attr("y", this.warehouse.y + size)
              .attr("width", size * 2)
              .attr("height", size * 2)
              .style("fill", "red");
            boxa
              .transition()
              .duration(duration)
              .attr("x", target.x + size)
              .attr("y", target.y + size)
              .on("end", () => boxa.remove());
          }
          if (key == "Baguette") {
            const boxb = this.svg
              .append("rect")
              .attr("x", this.warehouse.x + size)
              .attr("y", this.warehouse.y - size)
              .attr("width", size * 2)
              .attr("height", size * 2)
              .style("fill", "blue");
            boxb
              .transition()
              .duration(duration)
              .attr("x", target.x + size)
              .attr("y", target.y - size)
              .on("end", () => boxb.remove());
          }
          if (key == "Manchego Cheese") {
            const boxc = this.svg
              .append("rect")
              .attr("x", this.warehouse.x - size)
              .attr("y", this.warehouse.y - size)
              .attr("width", size * 2)
              .attr("height", size * 2)
              .style("fill", "green");
            boxc
              .transition()
              .duration(duration)
              .attr("x", target.x - size)
              .attr("y", target.y - size)
              .on("end", () => boxc.remove());
          }
          if (key == "Black Tea") {
            const boxd = this.svg
              .append("rect")
              .attr("x", this.warehouse.x - size)
              .attr("y", this.warehouse.y + size)
              .attr("width", size * 2)
              .attr("height", size * 2)
              .style("fill", "orange");
            boxd
              .transition()
              .duration(duration)
              .attr("x", target.x - size)
              .attr("y", target.y + size)
              .on("end", () => boxd.remove());
          }
        }
      });
    },

    //////////
    outlet1start() {
      this.outlet1 = new WebSocket("ws://localhost:8080/outlet1");
      this.outlet1.onmessage = (event) => {
        const data = JSON.parse(event.data);
        this.sendBoxToSupermarket(data);
      };
      this.outlet1.onclose = () => {
        console.log("WebSocket outlet1 Connection Closed");
      };
      this.outlet1.onerror = (error) => {
        console.error("WebSocket outlet1 Error:", error);
      };
    },

    outlet2start() {
      this.outlet2 = new WebSocket("ws://localhost:8080/outlet2");
      this.outlet2.onmessage = (event) => {
        const data = JSON.parse(event.data);
        this.sendBoxToSupermarket(data);
      };
      this.outlet2.onclose = () => {
        console.log("WebSocket outlet2 Connection Closed");
      };
      this.outlet2.onerror = (error) => {
        console.error("WebSocket outlet2 Error:", error);
      };
    },

    outlet3start() {
      this.outlet3 = new WebSocket("ws://localhost:8080/outlet3");
      this.outlet3.onmessage = (event) => {
        const data = JSON.parse(event.data);
        this.sendBoxToSupermarket(data);
      };
      this.outlet3.onclose = () => {
        console.log("WebSocket outlet3 Connection Closed");
      };
      this.outlet3.onerror = (error) => {
        console.error("WebSocket outlet3 Error:", error);
      };
    },

    outlet4start() {
      this.outlet4 = new WebSocket("ws://localhost:8080/outlet4");
      this.outlet4.onmessage = (event) => {
        const data = JSON.parse(event.data);
        this.sendBoxToSupermarket(data);
      };
      this.outlet4.onclose = () => {
        console.log("WebSocket outlet4 Connection Closed");
      };
      this.outlet4.onerror = (error) => {
        console.error("WebSocket outlet4 Error:", error);
      };
    },
    //////////
    toggleMessages1(messageBox) {
      if (messageBox === 1) {
        this.showMessages1 = !this.showMessages1;
      }
    },
    toggleMessages2(messageBox) {
      if (messageBox === 1) {
        this.showMessages2 = !this.showMessages2;
      }
    },
    toggleMessages3(messageBox) {
      if (messageBox === 1) {
        this.showMessages3 = !this.showMessages3;
      }
    },
    toggleMessages4(messageBox) {
      if (messageBox === 1) {
        this.showMessages4 = !this.showMessages4;
      }
    },

    centralMessages1(messageBox) {
      if (messageBox === 1) {
        this.showMessages5 = !this.showMessages5;
      }
    },
    centralMessages2(messageBox) {
      if (messageBox === 1) {
        this.showMessages6 = !this.showMessages6;
      }
    },
    centralMessages3(messageBox) {
      if (messageBox === 1) {
        this.showMessages7 = !this.showMessages7;
      }
    },
    centralMessages4(messageBox) {
      if (messageBox === 1) {
        this.showMessages8 = !this.showMessages8;
      }
    },

    getMessageClass(speakerid) {
      // 根据speakerid返回不同的CSS类名
      if (speakerid == "1") {
        return "speaker1-class";
      } else if (speakerid == "2") {
        return "speaker2-class";
      } else if (speakerid == "3") {
        return "speaker3-class";
      } else if (speakerid == "4") {
        return "speaker4-class";
      } else if (speakerid == "0") {
        return "speaker0-class"; // warehouse
      }
    },
    message1start() {
      this.message1 = new WebSocket("ws://localhost:8080/message1");
      this.message1.onmessage = (event) => {
        const data = JSON.parse(event.data);
        if (data.SpeakerID == "1") {
          this.messages1_text.push({
            id: this.nextMsgId1++,
            text: data.text,
            speakerid: data.SpeakerID,
          });
          //this.scrollToBottom();
          if (this.showMessages1) {
            const container1 = this.$refs.messages1;
            container1.scrollTop = container1.scrollHeight;
          }
        } else {
          //console.log(data.speakerid);
        }
      };
      this.message1.onopen = () => {
        console.log("WebSocket message1 connection established");
      };
      this.message1.onerror = (error) => {
        console.error("WebSocket message1 error observed:", error);
      };
      this.message1.onclose = () => {
        console.log("WebSocket message1 connection closed");
        this.message1 = null;
      };
    },
    message2start() {
      this.message2 = new WebSocket("ws://localhost:8080/message2");
      this.message2.onmessage = (event) => {
        const data = JSON.parse(event.data);
        if (data.SpeakerID == "2") {
          this.messages2_text.push({
            id: this.nextMsgId2++,
            text: data.text,
            speakerid: data.SpeakerID,
          });
          //this.scrollToBottom();
          if (this.showMessages2) {
            const container2 = this.$refs.messages2;
            container2.scrollTop = container2.scrollHeight;
          }
        } else {
          //console.log(data.speakerid);
        }
      };
      this.message2.onopen = () => {
        console.log("WebSocket message2 connection established");
      };
      this.message2.onerror = (error) => {
        console.error("WebSocket message2 error observed:", error);
      };
      this.message2.onclose = () => {
        console.log("WebSocket message2 connection closed");
        this.message2 = null;
      };
    },
    message3start() {
      this.message3 = new WebSocket("ws://localhost:8080/message3");
      this.message3.onmessage = (event) => {
        const data = JSON.parse(event.data);
        if (data.SpeakerID == "3") {
          this.messages3_text.push({
            id: this.nextMsgId3++,
            text: data.text,
            speakerid: data.SpeakerID,
          });
          if (this.showMessages3) {
            const container3 = this.$refs.messages3;
            container3.scrollTop = container3.scrollHeight;
          }
        } else {
          //console.log(data.speakerid);
        }
      };
      this.message3.onopen = () => {
        console.log("WebSocket message3 connection established");
      };
      this.message3.onerror = (error) => {
        console.error("WebSocket message3 error observed:", error);
      };
      this.message3.onclose = () => {
        console.log("WebSocket message3 connection closed");
        this.message3 = null;
      };
    },
    message4start() {
      this.message4 = new WebSocket("ws://localhost:8080/message4");
      this.message4.onmessage = (event) => {
        const data = JSON.parse(event.data);
        if (data.SpeakerID == "4") {
          this.messages4_text.push({
            id: this.nextMsgId4++,
            text: data.text,
            speakerid: data.SpeakerID,
          });
          if (this.showMessages4) {
            const container4 = this.$refs.messages4;
            container4.scrollTop = container4.scrollHeight;
          }
        } else {
          //console.log(data.speakerid);
        }
      };
      this.message4.onopen = () => {
        console.log("WebSocket message4 connection established");
      };
      this.message4.onerror = (error) => {
        console.error("WebSocket message4 error observed:", error);
      };
      this.message4.onclose = () => {
        console.log("WebSocket message4 connection closed");
        this.message4 = null;
      };
    },
    message5start() {
      this.message5 = new WebSocket("ws://localhost:8080/message5");
      this.message5.onmessage = (event) => {
        const data = JSON.parse(event.data);
        if (data.ReceiverID == "1") {
          this.messages1_text.push({
            id: this.nextMsgId1++,
            text: data.text,
            speakerid: data.SpeakerID,
          });
          //this.scrollToBottom();
          if (this.showMessages1) {
            const container1 = this.$refs.messages1;
            container1.scrollTop = container1.scrollHeight;
          }
        } else {
          //console.log(data.speakerid);
        }
      };
      this.message5.onopen = () => {
        console.log("WebSocket message5 connection established");
      };
      this.message5.onerror = (error) => {
        console.error("WebSocket message5 error observed:", error);
      };
      this.message5.onclose = () => {
        console.log("WebSocket message5 connection closed");
        this.message5 = null;
      };
    },
    message6start() {
      this.message6 = new WebSocket("ws://localhost:8080/message6");
      this.message6.onmessage = (event) => {
        const data = JSON.parse(event.data);
        if (data.ReceiverID == "2") {
          this.messages2_text.push({
            id: this.nextMsgId2++,
            text: data.text,
            speakerid: data.SpeakerID,
          });
          //this.scrollToBottom();
          if (this.showMessages2) {
            const container2 = this.$refs.messages2;
            container2.scrollTop = container2.scrollHeight;
          }
        } else {
          //console.log(data.speakerid);
        }
      };
      this.message6.onopen = () => {
        console.log("WebSocket message6 connection established");
      };
      this.message6.onerror = (error) => {
        console.error("WebSocket message6 error observed:", error);
      };
      this.message6.onclose = () => {
        console.log("WebSocket message6 connection closed");
        this.message6 = null;
      };
    },
    message7start() {
      this.message7 = new WebSocket("ws://localhost:8080/message7");
      this.message7.onmessage = (event) => {
        const data = JSON.parse(event.data);
        if (data.ReceiverID == "3") {
          this.messages3_text.push({
            id: this.nextMsgId3++,
            text: data.text,
            speakerid: data.SpeakerID,
          });
          if (this.showMessages3) {
            const container3 = this.$refs.messages3;
            container3.scrollTop = container3.scrollHeight;
          }
        } else {
          //console.log(data.speakerid);
        }
      };
      this.message7.onopen = () => {
        console.log("WebSocket message7 connection established");
      };
      this.message7.onerror = (error) => {
        console.error("WebSocket message7 error observed:", error);
      };
      this.message7.onclose = () => {
        console.log("WebSocket message7 connection closed");
        this.message7 = null;
      };
    },
    message8start() {
      this.message8 = new WebSocket("ws://localhost:8080/message8");
      this.message8.onmessage = (event) => {
        const data = JSON.parse(event.data);
        if (data.ReceiverID == "4") {
          this.messages4_text.push({
            id: this.nextMsgId4++,
            text: data.text,
            speakerid: data.SpeakerID,
          });
          if (this.showMessages4) {
            const container4 = this.$refs.messages4;
            container4.scrollTop = container4.scrollHeight;
          }
        } else {
          //console.log(data.speakerid);
        }
      };
      this.message8.onopen = () => {
        console.log("WebSocket message8 connection established");
      };
      this.message8.onerror = (error) => {
        console.error("WebSocket message8 error observed:", error);
      };
      this.message8.onclose = () => {
        console.log("WebSocket message8 connection closed");
        this.message8 = null;
      };
    },
  },
};
</script>

<style>
/* background #def6f6 */
/* table border #dc3636 */

#app {
  font-family: "Segoe UI", Tahoma, Geneva, Verdana, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #333;
  background-color: #def6f6;
  margin: 0;
  padding: 0;
}

header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  background-color: rgb(255, 255, 255);
  color: black;
  padding: 10px;
}

.utc {
  height: 100px;
}

.title {
  flex-grow: 1;
  text-align: center;
  font-size: 48px;
}

/* /////////////////////////////////////////////////// */

.btn,
.bth:link,
.btn:visited {
  display: inline-block;
  text-decoration: none;
  font-weight: 600;
  font-size: 28px;
  padding: 16px;
  border-radius: 9px;
  transition: all 300ms;
  cursor: pointer;
  border: none;
  font-family: inherit;
  width: 200px;
  height: 80px;
  background-color: rgb(241, 131, 131);
  color: white;
}

.btn:hover {
  background-color: white;
  color: red;
}

/* /////////////////////////////////////////////////// */

.information {
  padding: 20px;
  background: linear-gradient(to bottom, white, #def6f6);
}

.date {
  margin: 0 20px;
  font-size: 42px;
  color: #2e75c2;
  padding-bottom: 20px;
}

/* /////////////////////////////////////////////////// */

.event {
  font-size: 36px;
  color: #b35a00;
  font-weight: 600;
}

.event1 {
  position: absolute;
  top: 25px;
  left: 475px;
  width: 800px;
}

.event2 {
  position: absolute;
  top: 900px;
  left: 475px;
  width: 800px;
}

.event3 {
  position: absolute;
  top: 25px;
  left: 850px;
  width: 800px;
  text-align: right;
}

.event4 {
  position: absolute;
  top: 900px;
  left: 850px;
  width: 800px;
  text-align: right;
}

/* ///////////////////////////////////////////// */

.table {
  display: grid;
  grid-template-columns: repeat(5, 1fr);
}

.table th,
.table td {
  border: 1px solid #dc3636;
  color: black;
}

.table th {
  padding: 12px;
  background-color: #f2f2f2;
  font-size: 20px;
}

.table td {
  padding: 8px;
  font-weight: 500;
  text-align: left;
}

.warehousetable,
.supermarkettable {
  justify-self: center;
}

.table th:nth-child(2),
.table td:nth-child(2) {
  min-width: 50px;
}

.table tr:nth-child(1) td:nth-child(2) {
  background: linear-gradient(to right, #def6f6, 70%, red);
}
.table tr:nth-child(2) td:nth-child(2) {
  background: linear-gradient(to right, #def6f6, 70%, blue);
}
.table tr:nth-child(3) td:nth-child(2) {
  background: linear-gradient(to right, #def6f6, 70%, green);
}
.table tr:nth-child(4) td:nth-child(2) {
  background: linear-gradient(to right, #def6f6, 70%, orange);
}

/* ///////////////////////////////////////////// */

.communication {
  position: relative;
  height: 10px;
  width: 10px;
}

.communication {
  display: flex;
  justify-content: space-around;
  align-items: center;
  flex-wrap: wrap;
  margin-top: 20px;
  margin-left: 200px;
}

/* 添加下拉折叠按钮的样式 */
.toggle-button,
.bttoggle-buttonh:link,
.toggle-button:visited {
  display: inline-block;
  text-decoration: none;
  font-weight: 600;
  font-size: 25px;
  padding: 10px;
  border-radius: 9px;
  transition: all 300ms;
  cursor: pointer;
  border: none;
  font-family: inherit;
  width: 250px;
  height: 60px;
  background-color: rgb(241, 163, 18);
  color: white;
}

.toggle-button:hover {
  background-color: #def6f6;
  color: red;
}

.central-button,
.btcentral-button:link,
.central-button:visited {
  display: inline-block;
  text-decoration: none;
  font-weight: 600;
  font-size: 16px;
  padding: 10px;
  border-radius: 9px;
  transition: all 300ms;
  cursor: pointer;
  border: none;
  font-family: inherit;
  width: 180px;
  height: 40px;
  background-color: rgb(227, 54, 54);
  color: white;
}

.central-button:hover {
  background-color: white;
  color: red;
}

.show1 {
  position: absolute;
  top: 50px;
  left: -70px;
}

.show2 {
  position: absolute;
  top: 920px;
  left: -70px;
}

.show3 {
  position: absolute;
  top: 50px;
  left: 1900px;
}

.show4 {
  position: absolute;
  top: 920px;
  left: 1900px;
}

.centralshow1 {
  position: absolute;
  top: 320px;
  left: 550px;
}

.centralshow2 {
  position: absolute;
  top: 670px;
  left: 550px;
}

.centralshow3 {
  position: absolute;
  top: 320px;
  left: 850px;
}

.centralshow4 {
  position: absolute;
  top: 670px;
  left: 850px;
}

.message-container1 {
  position: absolute;
  top: 120px;
  left: -150px;
}

.message-container2 {
  position: absolute;
  top: 530px;
  left: -150px;
}

.message-container3 {
  position: absolute;
  top: 120px;
  left: 1900px;
}

.message-container4 {
  position: absolute;
  top: 530px;
  left: 1900px;
}

.message-container5 {
  position: absolute;
  top: 100px;
  left: 480px;
}

.message-container6 {
  position: absolute;
  top: 730px;
  left: 480px;
}

.message-container7 {
  position: absolute;
  top: 100px;
  left: 800px;
}

.message-container8 {
  position: absolute;
  top: 730px;
  left: 800px;
}

/* .message-container1,
.message-container2,
.message-container3,
.message-container4 {
  height: 180px;
  width: 280px;
  overflow-y: auto;
  border: 5px solid rgba(233, 201, 18, 0.867);
  border-radius: 12px;
  word-wrap: break-word;
  white-space: pre-wrap;
} */
/* .message-container1,
.message-container2,
.message-container3,
.message-container4 {
  border: 2px solid #ddd;
  border-radius: 4px;
  padding: 10px;
  margin: 10px 0;
  background-color: #f9f9f9;
  width: 300px;
  height: 330px;
  overflow-y: auto;
  word-wrap: break-word;
  white-space: pre-wrap;
} */
.message-container1,
.message-container2,
.message-container3,
.message-container4 {
  border: 1px solid #ddd;
  border-radius: 15px;
  padding: 10px;
  margin: 10px 0;
  background-color: #f8f5f5;
  /* 白色背景更接近聊天应用 */
  box-shadow: 0 10px 15px rgba(0, 0, 0, 0.1);
  /* 轻微的阴影效果 */
  width: 300px;
  /* 或根据您的布局调整 */
  height: 330px;
  overflow-y: auto;
  font-size: 20px;
  /* 更适合聊天的文字大小 */
  white-space: pre-wrap;
  word-wrap: break-word;
}

/* 聊天气泡样式 */
/* .message-container1 .message { */
/* background-color: #ceced3; */
/* 浅灰色聊天气泡，发送方可以选择不同颜色 */
/* border-radius: 18px; */
/* 圆润的聊天气泡角度 */
/* padding: 10px 15px; */
/* 内部填充 */
/* margin: 0px 0; */
/* 消息之间的间距 */
/* max-width: 80%; */
/* 气泡最大宽度 */
/* word-wrap: break-word; */
/* 自动换行 */
/* display: inline-block; */
/* 使元素像块级元素一样流动，但仍然是行内的 */
/* } */

/* 发送方的聊天气泡，您可以根据消息来源添加额外的类或ID */
.message-container1 .message.speaker1-class,
.message-container2 .message.speaker2-class,
.message-container3 .message.speaker3-class,
.message-container4 .message.speaker4-class {
  background-color: #b851ef;
  /* 蓝色气泡表示发送方 */
  color: white;
  /* 发送方气泡内的文字颜色 */
  /* margin-left: auto;
  margin-right: auto; */
  /* 发送方气泡靠右 */
  /* display: flex;
  flex-direction: column;
  align-items: flex-end; */
  /* 对齐到右侧 */
}

/* 接收方的聊天气泡 */
.message-container1 .message.speaker0-class,
.message-container2 .message.speaker0-class,
.message-container3 .message.speaker0-class,
.message-container4 .message.speaker0-class {
  background-color: #f1f0f0cf;
  color: #333;
}

.message-container5,
.message-container6,
.message-container7,
.message-container8 {
  height: 180px;
  width: 280px;
  overflow-y: auto;
  border: 5px solid rgba(198, 82, 80, 0.867);
  border-radius: 12px;
  word-wrap: break-word;
  /* 允许长单词或无间断的文本换行 */
  white-space: pre-wrap;
}

.message-container1 span,
.message-container2 span,
.message-container3 span,
.message-container4 span,
.message-container5 span,
.message-container6 span,
.message-container7 span,
.message-container8 span {
  display: inline;
  /* 使用行内元素 */
}

/* .speaker1-class {
  color: rgb(255, 0, 0);
}

.speaker2-class {
  color: rgb(0, 42, 255);
}

.speaker3-class {
  color: rgb(1, 1, 19);
}

.speaker4-class {
  color: rgb(157, 0, 255);
}

.speaker0-class {
  color: rgb(255, 136, 0);
} */
.speaker1-class {
  color: #e91e63;
}

.speaker2-class {
  color: #03a9f4;
}

.speaker3-class {
  color: #4caf50;
}

.speaker4-class {
  color: #ffc107;
}

.speaker0-class {
  color: #9c27b0;
}
</style>
