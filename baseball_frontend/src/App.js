import React, { Component } from "react";
import axios from "axios";

import logo from "./logo.svg";
import "./App.css";
import baseball from "./baseball.png";

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
      .get("http://localhost:8081/", {
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

class BaseballTable extends Component {
  state = {
    rows: this.props.teams.map(team => {
      team.expanded = false;
      return team;
    })
  };

  render() {
    const { rows } = this.state;

    return (
      <div>
        <img src={baseball} alt="dank_baseball" />
        <h1>
          <span>⚾</span> Baseball <span> ⚾</span>
        </h1>
        <table className="highlight bordered">
          <thead>
            <tr>
              <th>Rank</th>
              <th>Owner</th>
              <th>Name</th>
              <th>Wins</th>
              <th>Losses</th>
              <th>Pct</th>
            </tr>
          </thead>
          <tbody>
            {rows.map(team => (
              <React.Fragment>
                <tr onClick={() => this.handleRowClick(team)}>
                  <td> {team.rank} </td>
                  <td> {team.owner} </td>
                  <td> {team.name} </td>
                  <td> {team.wins} </td>
                  <td> {team.losses} </td>
                  <td> {team.perc} </td>
                </tr>
                {team.expanded &&
                  team.teams.map(team => (
                    <tr className="expanded-row">
                      <td>
                        <div />
                      </td>
                      <td>
                        <div>{team.owner}</div>
                      </td>
                      <td>
                        <div>
                          {team.first_name} {team.last_name}
                        </div>
                      </td>
                      <td>
                        <div>{team.won} </div>
                      </td>
                      <td>
                        <div>{team.lost} </div>
                      </td>
                      <td>
                        <div>{team.win_percentage} </div>
                      </td>
                    </tr>
                  ))}
              </React.Fragment>
            ))}
          </tbody>
        </table>
      </div>
    );
  }

  handleRowClick(team) {
    console.log({ team });
    const { rows } = this.state;
    const newRows = [...rows];
    const expandedRow = newRows.find(row => row.name === team.name);
    expandedRow.expanded = !expandedRow.expanded;
    this.setState({
      rows: newRows
    });
  }
}

export default App;
