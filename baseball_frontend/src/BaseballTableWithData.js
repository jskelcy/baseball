import React, { Component } from "react";
import axios from "axios";

import "./App.css";
import BaseballTable from "./BaseballTable";
import AppRouter from "./Routes";
import { BaseballError } from "./Error";

class BaseballTableWithData extends Component {
  state = {
    loading: true,
    error: false,
    teams: []
  };

  async componentDidMount() {
    // TODO: this should probabl get
    // configured via environment vars

    const { league } = this.props;

    let leagueEndpoint;
    switch (league) {
      case "cork":
        leagueEndpoint = "CORK";
        break;
      default:
        leagueEndpoint = "CHRIS";
        break;
    }

    const endpoint =
      process.env.NODE_ENV === "production"
        ? `http://${
            window.location.href.includes("www") ? "www." : ""
          }baseballis.cool/`
        : "http://localhost:8081/";

    const response = await axios
      .get(`${endpoint}${leagueEndpoint}`, {
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

export default BaseballTableWithData;
