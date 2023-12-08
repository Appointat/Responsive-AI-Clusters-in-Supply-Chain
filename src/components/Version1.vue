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
        centralhub: null,
        svg: null,
        warehouse: { x: 800, y: 500 },
        supermarkets: [
          { x: 150, y: 150 },
          { x: 150, y: 850 },
          { x: 1450, y: 150 },
          { x: 1450, y: 850 },
        ],
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
          this.svg
            .append("image")
            .attr("xlink:href", require("@/assets/supermarket.png"))
            .attr("x", s.x - 15)
            .attr("y", s.y - 15)
            .attr("width", 55)
            .attr("height", 55);
          // supermarket number
          this.svg
            .append("text")
            .attr("x", s.x - 35)
            .attr("y", s.y - 20)
            .text(`${supermarketNames[index]}`)
            .style("font-size", "28px")
            .style("font-family", "Arial, sans-serif")
            .attr("class", "supermarket-label");
        });
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
  
      handlemessage(messagebag) {
        switch (messagebag.SpeakerID) {
          case "1":
            this.printmessage(messagebag.text, { x: 300, y: 100 });
            break;
          case "2":
            this.printmessage(messagebag.text, { x: 50, y: 650 });
            break;
          case "3":
            this.printmessage(messagebag.text, { x: 1250, y: 250 });
            break;
          case "4":
            this.printmessage(messagebag.text, { x: 1050, y: 750 });
            break;
          case "0":
            switch (messagebag.ReceiverID) {
              case "1":
                this.printmessage(messagebag.text, { x: 500, y: 300 });
                break;
              case "2":
                this.printmessage(messagebag.text, { x: 430, y: 490 });
                break;
              case "3":
                this.printmessage(messagebag.text, { x: 900, y: 360 });
                break;
              case "4":
                this.printmessage(messagebag.text, { x: 750, y: 600 });
            }
        }
      },
  
      printmessage(message, position) {
        const textPadding = 5;
        const maxWidth = 300;
        const lineHeight = 20;
        const words = message.split(/\s+/);
        let line = "";
        let lines = [];
  
        words.forEach((word) => {
          const testLine = line + (line === "" ? "" : " ") + word;
          const testWidth = this.getTextWidth(testLine);
  
          if (testWidth > maxWidth && line !== "") {
            lines.push(line);
            line = word;
          } else {
            line = testLine;
          }
        });
  
        lines.push(line);
  
        const bubbleHeight = lines.length * lineHeight + textPadding * 2;
  
        const bubble = this.svg
          .append("rect")
          .attr("x", position.x)
          .attr("y", position.y)
          .attr("width", maxWidth)
          .attr("height", bubbleHeight)
          .attr("rx", 10)
          .attr("ry", 10)
          .style("fill", "white")
          .style("stroke", "black")
          .style("stroke-width", 1);
  
        // 添加多行文本
        const text = this.svg
          .selectAll(".text-line")
          .data(lines)
          .enter()
          .append("text")
          .attr("class", "text-line")
          .attr("x", position.x + textPadding)
          .attr("y", (d, i) => position.y + textPadding + i * lineHeight + 10)
          .text((d) => d);
        console.log("aaa");
        // 添加过渡效果和消失
        this.svg
          .selectAll(".text-line")
          .transition()
          .duration(3000)
          //.style("opacity", 1)
          .on("end", () => {
            bubble.remove();
            text.remove();
            d3.selectAll(".text-line").remove();
          });
        console.log("bbb");
      },
  
      getTextWidth(text) {
        // 返回文本宽度的辅助函数
        const dummyText = this.svg.append("text").text(text);
        const width = dummyText.node().getBBox().width;
        dummyText.remove();
        return width;
      },
  
      message1start() {
        this.message1 = new WebSocket("ws://localhost:8080/message1");
        this.message1.onmessage = (event) => {
          const data = JSON.parse(event.data);
          this.handlemessage(data);
        };
        this.message1.onclose = () => {
          console.log("WebSocket Connection Closed");
        };
        this.message1.onerror = (error) => {
          console.error("WebSocket Error:", error);
        };
      },
  
      message2start() {
        this.message2 = new WebSocket("ws://localhost:8080/message2");
        this.message2.onmessage = (event) => {
          const data = JSON.parse(event.data);
          this.handlemessage(data);
        };
        this.message2.onclose = () => {
          console.log("WebSocket Connection Closed");
        };
        this.message2.onerror = (error) => {
          console.error("WebSocket Error:", error);
        };
      },
  
      message3start() {
        this.message3 = new WebSocket("ws://localhost:8080/message3");
        this.message3.onmessage = (event) => {
          const data = JSON.parse(event.data);
          this.handlemessage(data);
        };
        this.message3.onclose = () => {
          console.log("WebSocket Connection Closed");
        };
        this.message3.onerror = (error) => {
          console.error("WebSocket Error:", error);
        };
      },
  
      message4start() {
        this.message4 = new WebSocket("ws://localhost:8080/message4");
        this.message4.onmessage = (event) => {
          const data = JSON.parse(event.data);
          this.handlemessage(data);
        };
        this.message4.onclose = () => {
          console.log("WebSocket Connection Closed");
        };
        this.message4.onerror = (error) => {
          console.error("WebSocket Error:", error);
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
  </style>
  