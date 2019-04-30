import React, { Component } from "react";
import axios from "axios";

import "./App.css";
import BaseballTable from "./baseball_table.js";

class App extends Component {
  state = {
    loading: true,
    error: false,
    teams: []
  };

  async componentDidMount() {
    // TODO: this should probabl get
    // configured via environment vars
    const response = await axios
      .get("http://localhost:8081/CHRIS", {
        headers: { "Access-Control-Allow-Origin": "*" }
      })
      .catch(error => {
        this.setState({
          loading: false,
          error: true
        });
      });

    if (response) {
      if (response.data.status !== 200) {
        return this.setState({
          loading: false,
          error: response.data.error
        });
      }
      return this.setState({ loading: false, teams: response.data.teams });
    }
  }

  render() {
    const { loading, teams, error } = this.state;
    // TODO: baseball loading spinner
    return (
      <div className="App">
        {error ? (
          <BaseballError error={error} />
        ) : loading ? (
          "loading sick baseball shit..."
        ) : (
          <BaseballTable teams={teams} />
        )}
      </div>
    );
  }
}

const BaseballError = ({ error }) => (
  <h1>
    <span className="App-logo">⚾⚾⚾ </span> {error} beisbol ¯\_(ツ)_/¯ 🚬{" "}
  </h1>
);

export default App;
