import React, { Component } from "react";
import axios from "axios";

import logo from "./logo.svg";
import "./App.css";

class App extends Component {
  state = {
    loading: false
  };

  async componentDidMount() {
    this.setState({ loading: true });
    // TODO: this should probabl get
    // configured via environment vars
    const data = await axios.get("http://localhost:8080/hello-world", {
      headers: { "Access-Control-Allow-Origin": "*" }
    });
    console.log({ data });
    this.setState({ loading: false });
  }

  render() {
    const { loading } = this.state;
    // TODO: baseball loading spinner
    return (
      <div className="App">
        {loading ? "loading sick baseball shit..." : <BaseballTable />}
      </div>
    );
  }
}

const BaseballTable = props => (
  <div>
    <h1>Baseball</h1>
    <table className="highlight bordered">
      <thead>
        <tr>
          <th>Rank</th>
          <th>Teams</th>
          <th>Owner</th>
          <th>Wins</th>
          <th>Losses</th>
          <th>Pct</th>
        </tr>
      </thead>
      <tbody>
        {/* {{range .League}}
        <tr>
        <td> {{.Rank}} </td>
        <td> {{.Name}} </td>
        <td> {{.Owner}} </td>
        <td> {{.Wins}} </td>
        <td> {{.Losses}} </td>
        <td> {{.RenderPerc}} </td>
        </tr>
    {{end}} */}
      </tbody>
    </table>
  </div>
);

export default App;
