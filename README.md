# Responsive AI Clusters in Supply Chain

## Introduction

Welcome to "Responsive AI Clusters in Supply Chain" - a groundbreaking project aimed at revolutionizing supply chain management through the use of responsive, intelligent multi-agent systems.

![Surface](image/surface.png)

![Responsive AI Clusters](https://github.com/Appointat/Responsive-AI-Clusters-in-Supply-Chain/assets/65004114/02731b24-c7ad-41b5-8779-cf935a65b919)

Our project mainly revolves around simulating a supply chain, where the entire supply chain system is roughly composed of a central warehouse and various offline outlets (mainly supermarkets). In reality, there may be multiple central warehouses, but in this program, for the sake of simplifying the simulation, we have implemented a singleton pattern for the central warehouse. All other outlets communicate daily with the central warehouse to receive restocks. Goods are a separate class containing information such as name and quantity. To simplify the simulation program, we have only instantiated four types of products: olive oil, baguette, manchego cheese, black tea.

In our program, we instantiated four outlets and one central warehouse. The central warehouse is an independent AI Agent, whose inventory changes autonomously based on certain conditions (for example, when stock levels are low, the AI might increase the inventory of certain products through other events). Each outlet has its own independent event chain, which is stored in the event.go file. When the date of an event (such as an unexpected incident, holiday, celebration, etc.) matches the current clock date, the event's description will be sent to the AI side. Since each outlet has a unique ID, the AI side identifies which outlet sent the message through the ID and allocates it to the corresponding AI agent.
## Project Background and Prospects

"In the actual application process, customers in the location of each outlet have a unified preference. In this project, we have directly assigned a basic description to the local customer preferences of each outlet. Of course, in reality, these preferences can change over time. Theoretically, this aspect could also be managed by AI. If there are multiple central warehouses, then each outlet needs to consider its own inventory status, the distance to each warehouse, among other factors, to choose the most optimal central warehouse for restocking and other operations.
## Motivation

The motivation behind this project stems from the increasing complexity of modern supply chains and the need for more dynamic, real-time decision-making processes. Traditional supply chain mechanisms are often static and can't adapt quickly to the ever-changing market demands or unforeseen disruptions. This project introduces a flexible, scalable solution that not only responds to current conditions but also anticipates future challenges, optimizing the supply chain for resilience and efficiency.

## Principles

Our approach is based on several key principles:

- **Multi-Agent Collaboration**: Harnessing the power of AI agents, each representing entities within the supply chain, enabling decentralized decision-making and fostering robust collaboration.
- **Real-Time Responsiveness**: Ensuring the system is capable of adapting to new events and information, maintaining supply chain continuity and efficiency.
- **Predictive Analytics**: Utilizing advanced data analytics to forecast demand and supply scenarios, allowing for preemptive strategy adjustments.
- **Scalability and Flexibility**: Designing the system to be inherently scalable, handling the expansion seamlessly and adapting to various supply chain sizes and structures.
- **Sustainability**: Focusing on long-term sustainability by optimizing resource allocation and reducing waste.

## System Architecture

[Insert a block diagram or flowchart]

The architecture of our system is structured around a central distribution hub, surrounded by retail outlets, each equipped with AI agents. These agents communicate with the central hub to balance supply with demand, share resources, and optimize the overall network performance.

## Commercial Value

The commercial implications of implementing such a system are vast:

- **Cost Reduction**: Through efficient resource allocation and waste minimization.
- **Increased Agility**: Enabling businesses to quickly adapt to market changes or disruptions.
- **Enhanced Decision-Making**: Data-driven insights allow for more informed and strategic business decisions.
- **Competitive Advantage**: Businesses equipped with responsive supply chains can outmaneuver competition and respond better to consumer needs.

Certainly, here's the completed Markdown documentation for the Backend: GPT-4-turbo installation and getting started process:

## Installation
Clone the repository. If you have already downloaded the whole project, you can skip this step.
```sh
git clone https://github.com/Appointat/Responsive-AI-Clusters-in-Supply-Chain.git
```
### Backend: GPT-4-turbo

1. Create a new terminal.
2. Navigate to the `backend/ai` directory.
   ```sh
   cd ./backend/ai
   ```
3. Activate the virtual environment.
   ```sh
   source venv/bin/activate  # On Unix or macOS
   venv\Scripts\activate     # On Windows
   ```
4. Configure the OpenAI API key.
   ```sh
   export OPENAI_API_KEY="your-openai-api-key"  # On Unix or macOS
   set OPENAI_API_KEY="your-openai-api-key"     # On Windows
   ```


### Backend: Go Routines

To install the Go Routines backend, follow these steps:
1. Create a new terminal.
2. Navigate to the `backend/go_routine` directory.
   ```sh
   cd ./backend/go_routine
   ```


### UI: Vue.js

To install the Vue.js frontend, follow these steps:
1. Create a new terminal.
2. Navigate to the `frontend` directory.
   ```sh
   cd ./frontend
   ```
3. Install the necessary modules.
   ```sh
   npm install
   ```


## Getting Started
### UI: Vue.js
1. Start the web frontend.
   ```sh
   npm run serve
   ```

### Backend: Go Routines
1. Run the Go application.
   ```sh
   go run main.go
   ```

### Backend: GPT-4-turbo
1. Run the application.
   ```sh
   python app.py
   ```

## Contribution

We are open to contributions! Please read through `CONTRIBUTING.md` for guidelines on how to make a pull request.

## License

This project is licensed under the [MIT License](LICENSE) - see the LICENSE file for details.

## Acknowledgments

[List of contributors and acknowledgments]

We thank everyone who contributes to the development and advancement of this project.

## Contact

[Your Name] - [Your Email]

Project Link: [GitHub Project Link]

---

We invite you to join us in shaping the future of supply chain management with Responsive AI Clusters. Let's build a smarter, more adaptive future together.
