<template>
  <header>
    <img class="utc" alt="utc logo" src="@/assets/utc_logo.jpg" />
    <h1 class="title">Responsive AI Clusters in Supply Chain</h1>
    <button @click="start" class="btn">Start</button>
  </header>
  <div class="information">
    <h2 class="date">Date: {{ date }}</h2>
    <h2 class="event">Event: {{ event }}</h2>
  </div>
  <div ref="visualization">
    <div class="table">
      <div class="warehousetable">
        <table>
          <thead>
            <tr>
              <th>Warehouse</th>
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
              <th>Auchan</th>
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
              <th>Carrefour</th>
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
              <th>Monoprix</th>
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
              <th>Normal</th>
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
      <button @click="toggleMessages1(1)" class="toggle-button show1">Auchan Messages</button>
      <button @click="toggleMessages2(1)" class="toggle-button show2">Carrefour Messages</button>
      <button @click="toggleMessages3(1)" class="toggle-button show3">Monoprix Messages</button>
      <button @click="toggleMessages4(1)" class="toggle-button show4">Normal Messages</button>
      <button @click="centralMessages1(1)" class="central-button centralshow1">Central to Auchan</button>
      <button @click="centralMessages2(1)" class="central-button centralshow2">Central to Carrefour</button>
      <button @click="centralMessages3(1)" class="central-button centralshow3">Central to Monoprix</button>
      <button @click="centralMessages4(1)" class="central-button centralshow4">Central to Normal</button>
      <!-- supermarkets messages -->
      <div ref="messages1" class="message-container1" v-if="message1 && showMessages1">
        <span v-for="msg in messages1_text" :key="msg.id">{{ msg.text }} </span>
      </div>
      <div ref="messages2" class="message-container2" v-if="message2 && showMessages2">
        <span v-for="msg in messages2_text" :key="msg.id">{{ msg.text }} </span>
      </div>
      <div ref="messages3" class="message-container3" v-if="message3 && showMessages3">
        <span v-for="msg in messages3_text" :key="msg.id">{{ msg.text }} </span>
      </div>
      <div ref="messages4" class="message-container4" v-if="message4 && showMessages4">
        <span v-for="msg in messages4_text" :key="msg.id">{{ msg.text }} </span>
      </div>
      <!-- warehouse messages -->
      <div ref="messages5" class="message-container5" v-if="message5 && showMessages5">
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
      </div>
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
        ["A", 0],
        ["B", 0],
        ["C", 0],
        ["D", 0],
      ]),
      stock1: new Map([
        ["A", 0],
        ["B", 0],
        ["C", 0],
        ["D", 0],
      ]),
      stock2: new Map([
        ["A", 0],
        ["B", 0],
        ["C", 0],
        ["D", 0],
      ]),
      stock3: new Map([
        ["A", 0],
        ["B", 0],
        ["C", 0],
        ["D", 0],
      ]),
      stock4: new Map([
        ["A", 0],
        ["B", 0],
        ["C", 0],
        ["D", 0],
      ]),
      onedaytime: 5000,
      date: null,
      event: null,
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
      warehouse: { x: 800, y: 500 },
      supermarkets: [
        { x: 150, y: 150 },
        { x: 150, y: 850 },
        { x: 1450, y: 150 },
        { x: 1450, y: 850 },
      ],
      messages1_text: [],
      messages2_text: [],
      messages3_text: [],
      messages4_text: [],
      messages5_text: [],
      messages6_text: [],
      messages7_text: [],
      messages8_text: [],
      nextMsgId1: 0,
      nextMsgId2: 0,
      nextMsgId3: 0,
      nextMsgId4: 0,
      nextMsgId5: 0,
      nextMsgId6: 0,
      nextMsgId7: 0,
      nextMsgId8: 0,
      showMessages1: true,
      showMessages2: true,
      showMessages3: true,
      showMessages4: true,
      showMessages5: true,
      showMessages6: true,
      showMessages7: true,
      showMessages8: true
    };
  },
  //////////
  mounted() {
    this.createVisualization();
  },
  //////////
  methods: {
    createVisualization() {
      const width = 1600;
      const height = 1000;
      // Create SVG canvas
      this.svg = d3
        .select(this.$refs.visualization)
        .append("svg")
        .attr("width", width)
        .attr("height", height);
      // Drawing the line between warehouse and supermarket
      this.supermarkets.forEach((s) => {
        this.svg
          .append("line")
          .style("stroke", "black")
          .attr("x1", this.warehouse.x)
          .attr("y1", this.warehouse.y)
          .attr("x2", s.x)
          .attr("y2", s.y + 15);
      });
      // warehouse
      this.svg
        .append("image")
        .attr("xlink:href", require("@/assets/warehouse.png"))
        .attr("x", this.warehouse.x - 30)
        .attr("y", this.warehouse.y - 30)
        .attr("width", 60)
        .attr("height", 60);
      // supermarket
      const supermarketNames = ["Auchan", "Carrefour", "Monoprix", "Normal"];
      this.supermarkets.forEach((s, index) => {
        //this.svg
        //.append("image")
        //.attr("xlink:href", require("@/assets/supermarket.png"))
        //.attr("x", s.x - 15)
        //.attr("y", s.y - 15)
        //.attr("width", 55)
        //.attr("height", 55);
        // supermarket number
        this.svg
          .append("text")
          .attr("x", s.x - 30)
          .attr("y", s.y - 20)
          .text(`${supermarketNames[index]}`)
          .style("font-size", "28px")
          .style("font-family", "Arial, sans-serif")
          .attr("class", "supermarket-label");
      });
      this.svg
        .append("image")
        .attr("xlink:href", require("@/assets/auchan.jpg"))
        .attr("x", this.supermarkets[0].x - 15)
        .attr("y", this.supermarkets[0].y - 15)
        .attr("width", 80)
        .attr("height", 80);
      this.svg
        .append("image")
        .attr("xlink:href", require("@/assets/carrefour.jpg"))
        .attr("x", this.supermarkets[1].x - 15)
        .attr("y", this.supermarkets[1].y - 15)
        .attr("width", 80)
        .attr("height", 80);
      this.svg
        .append("image")
        .attr("xlink:href", require("@/assets/monoprix.jpg"))
        .attr("x", this.supermarkets[2].x - 15)
        .attr("y", this.supermarkets[2].y - 15)
        .attr("width", 80)
        .attr("height", 80);
      this.svg
        .append("image")
        .attr("xlink:href", require("@/assets/normal.jpg"))
        .attr("x", this.supermarkets[3].x - 15)
        .attr("y", this.supermarkets[3].y - 15)
        .attr("width", 80)
        .attr("height", 80);
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
      this.event = generalInfo.event;
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
      //////////
      Object.entries(supermarketInfo.productAdd).forEach((table) => {
        let key = table[0];
        let value = table[1];
        if (value != 0) {
          if (key == "A") {
            console.log("321");
            const boxa = this.svg
              .append("rect")
              .attr("x", this.warehouse.x + 5)
              .attr("y", this.warehouse.y + 5)
              .attr("width", 10)
              .attr("height", 10)
              .style("fill", "red");
            boxa
              .transition()
              .duration(duration)
              .attr("x", target.x + 5)
              .attr("y", target.y + 5)
              .on("end", () => boxa.remove());
          }
          if (key == "B") {
            const boxb = this.svg
              .append("rect")
              .attr("x", this.warehouse.x + 5)
              .attr("y", this.warehouse.y - 5)
              .attr("width", 10)
              .attr("height", 10)
              .style("fill", "blue");
            boxb
              .transition()
              .duration(duration)
              .attr("x", target.x + 5)
              .attr("y", target.y - 5)
              .on("end", () => boxb.remove());
          }
          if (key == "C") {
            const boxc = this.svg
              .append("rect")
              .attr("x", this.warehouse.x - 5)
              .attr("y", this.warehouse.y - 5)
              .attr("width", 10)
              .attr("height", 10)
              .style("fill", "green");
            boxc
              .transition()
              .duration(duration)
              .attr("x", target.x - 5)
              .attr("y", target.y - 5)
              .on("end", () => boxc.remove());
          }
          if (key == "D") {
            const boxd = this.svg
              .append("rect")
              .attr("x", this.warehouse.x - 5)
              .attr("y", this.warehouse.y + 5)
              .attr("width", 10)
              .attr("height", 10)
              .style("fill", "orange");
            boxd
              .transition()
              .duration(duration)
              .attr("x", target.x - 5)
              .attr("y", target.y + 5)
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
        console.log("WebSocket Connection Closed");
      };
      this.outlet1.onerror = (error) => {
        console.error("WebSocket Error:", error);
      };
    },

    outlet2start() {
      this.outlet2 = new WebSocket("ws://localhost:8080/outlet2");
      this.outlet2.onmessage = (event) => {
        const data = JSON.parse(event.data);
        this.sendBoxToSupermarket(data);
      };
      this.outlet2.onclose = () => {
        console.log("WebSocket Connection Closed");
      };
      this.outlet2.onerror = (error) => {
        console.error("WebSocket Error:", error);
      };
    },

    outlet3start() {
      this.outlet3 = new WebSocket("ws://localhost:8080/outlet3");
      this.outlet3.onmessage = (event) => {
        const data = JSON.parse(event.data);
        this.sendBoxToSupermarket(data);
      };
      this.outlet3.onclose = () => {
        console.log("WebSocket Connection Closed");
      };
      this.outlet3.onerror = (error) => {
        console.error("WebSocket Error:", error);
      };
    },

    outlet4start() {
      this.outlet4 = new WebSocket("ws://localhost:8080/outlet4");
      this.outlet4.onmessage = (event) => {
        const data = JSON.parse(event.data);
        this.sendBoxToSupermarket(data);
      };
      this.outlet4.onclose = () => {
        console.log("WebSocket Connection Closed");
      };
      this.outlet4.onerror = (error) => {
        console.error("WebSocket Error:", error);
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


    message1start() {
      this.message1 = new WebSocket('ws://localhost:8080/message1');
      this.message1.onmessage = (event) => {
        const data = JSON.parse(event.data);
        if (data.SpeakerID == "1") {
          this.messages1_text.push({ id: this.nextMsgId1++, text: data.text });
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
        console.log("WebSocket connection established");
      };
      this.message1.onerror = (error) => {
        console.error("WebSocket error observed:", error);
      };
      this.message1.onclose = () => {
        console.log("WebSocket connection closed");
        this.message1 = null;
      };
    },
    message2start() {
      this.message2 = new WebSocket('ws://localhost:8080/message2');
      this.message2.onmessage = (event) => {
        const data = JSON.parse(event.data);
        if (data.SpeakerID == "2") {
          this.messages2_text.push({ id: this.nextMsgId2++, text: data.text });
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
        console.log("WebSocket connection established");
      };
      this.message2.onerror = (error) => {
        console.error("WebSocket error observed:", error);
      };
      this.message2.onclose = () => {
        console.log("WebSocket connection closed");
        this.message2 = null;
      };
    },
    message3start() {
      this.message3 = new WebSocket('ws://localhost:8080/message3');
      this.message3.onmessage = (event) => {
        const data = JSON.parse(event.data);
        if (data.SpeakerID == "3") {
          this.messages3_text.push({ id: this.nextMsgId3++, text: data.text });
          if (this.showMessages3) {
            const container3 = this.$refs.messages3;
            container3.scrollTop = container3.scrollHeight;
          }
        } else {
          //console.log(data.speakerid);
        }
      };
      this.message3.onopen = () => {
        console.log("WebSocket connection established");
      };
      this.message3.onerror = (error) => {
        console.error("WebSocket error observed:", error);
      };
      this.message3.onclose = () => {
        console.log("WebSocket connection closed");
        this.message3 = null;
      };
    },
    message4start() {
      this.message4 = new WebSocket('ws://localhost:8080/message4');
      this.message4.onmessage = (event) => {
        const data = JSON.parse(event.data);
        if (data.SpeakerID == "4") {
          this.messages4_text.push({ id: this.nextMsgId4++, text: data.text });
          if (this.showMessages4) {
            const container4 = this.$refs.messages4;
            container4.scrollTop = container4.scrollHeight;
          }
        } else {
          //console.log(data.speakerid);
        }
      };
      this.message4.onopen = () => {
        console.log("WebSocket connection established");
      };
      this.message4.onerror = (error) => {
        console.error("WebSocket error observed:", error);
      };
      this.message4.onclose = () => {
        console.log("WebSocket connection closed");
        this.message4 = null;
      };
    },
    message5start() {
      this.message5 = new WebSocket('ws://localhost:8080/message5');
      this.message5.onmessage = (event) => {
        const data = JSON.parse(event.data);
        if (data.ReceiverID == "1") {
          this.messages5_text.push({ id: this.nextMsgId5++, text: data.text });
          //this.scrollToBottom();
          if (this.showMessages5) {
            const container5 = this.$refs.messages5;
            container5.scrollTop = container5.scrollHeight;
          }
        } else {
          //console.log(data.speakerid);
        }
      };
      this.message5.onopen = () => {
        console.log("WebSocket connection established");
      };
      this.message5.onerror = (error) => {
        console.error("WebSocket error observed:", error);
      };
      this.message5.onclose = () => {
        console.log("WebSocket connection closed");
        this.message5 = null;
      };
    },
    message6start() {
      this.message6 = new WebSocket('ws://localhost:8080/message6');
      this.message6.onmessage = (event) => {
        const data = JSON.parse(event.data);
        if (data.ReceiverID == "2") {
          this.messages6_text.push({ id: this.nextMsgId6++, text: data.text });
          //this.scrollToBottom();
          if (this.showMessages6) {
            const container6 = this.$refs.messages6;
            container6.scrollTop = container6.scrollHeight;
          }
        } else {
          //console.log(data.speakerid);
        }
      };
      this.message6.onopen = () => {
        console.log("WebSocket connection established");
      };
      this.message6.onerror = (error) => {
        console.error("WebSocket error observed:", error);
      };
      this.message6.onclose = () => {
        console.log("WebSocket connection closed");
        this.message6 = null;
      };
    },
    message7start() {
      this.message7 = new WebSocket('ws://localhost:8080/message7');
      this.message7.onmessage = (event) => {
        const data = JSON.parse(event.data);
        if (data.ReceiverID == "3") {
          this.messages7_text.push({ id: this.nextMsgId7++, text: data.text });
          if (this.showMessages7) {
            const container7 = this.$refs.messages7;
            container7.scrollTop = container7.scrollHeight;
          }
        } else {
          //console.log(data.speakerid);
        }
      };
      this.message7.onopen = () => {
        console.log("WebSocket connection established");
      };
      this.message7.onerror = (error) => {
        console.error("WebSocket error observed:", error);
      };
      this.message7.onclose = () => {
        console.log("WebSocket connection closed");
        this.message7 = null;
      };
    },
    message8start() {
      this.message8 = new WebSocket('ws://localhost:8080/message8');
      this.message8.onmessage = (event) => {
        const data = JSON.parse(event.data);
        if (data.ReceiverID == "4") {
          this.messages8_text.push({ id: this.nextMsgId8++, text: data.text });
          if (this.showMessages8) {
            const container8 = this.$refs.messages8;
            container8.scrollTop = container8.scrollHeight;
          }
        } else {
          //console.log(data.speakerid);
        }
      };
      this.message8.onopen = () => {
        console.log("WebSocket connection established");
      };
      this.message8.onerror = (error) => {
        console.error("WebSocket error observed:", error);
      };
      this.message8.onclose = () => {
        console.log("WebSocket connection closed");
        this.message8 = null;
      };
    },
  },
};
</script>
  
<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  /* text-align: center; */
  color: #2c3e50;
}

header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.utc {
  height: 100px;
}

.title {
  flex-grow: 1;
  text-align: center;
  font-size: 48px;
}

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

.information {
  display: flex;
  justify-content: flex-start;
}

.date {
  flex: 1;
}

.event {
  flex: 2;
}

.table {
  grid-template-columns: repeat(5, 1fr);
  display: grid;
  justify-content: space-between;
}

.supermarkettable {
  justify-self: center;
}

.warehousetable {
  justify-self: center;
}

.communication {
  position: relative;
  height: 10px;
  width: 10px;
}

/* 添加下拉折叠按钮的样式 */
.toggle-button,
.bttoggle-buttonh:link,
.toggle-button:visited {
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
  background-color: rgb(222, 211, 0);
  color: white;
}

.toggle-button:hover {
  background-color: white;
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
  top: 220px;
  left: 50px;
}

.show2 {
  position: absolute;
  top: 770px;
  left: 50px;
}

.show3 {
  position: absolute;
  top: 220px;
  left: 1400px;
}

.show4 {
  position: absolute;
  top: 770px;
  left: 1400px;
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
  top: 280px;
  left: 50px;
}

.message-container2 {
  position: absolute;
  top: 570px;
  left: 50px;
}

.message-container3 {
  position: absolute;
  top: 280px;
  left: 1400px;
}

.message-container4 {
  position: absolute;
  top: 570px;
  left: 1400px;
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

.message-container1,
.message-container2,
.message-container3,
.message-container4 {
  height: 180px;
  width: 280px;
  overflow-y: auto;
  border: 5px solid rgba(233, 201, 18, 0.867);
  border-radius: 12px;
  word-wrap: break-word;
  /* 允许长单词或无间断的文本换行 */
  white-space: pre-wrap;
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
</style>
  