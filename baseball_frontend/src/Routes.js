import React from "react";
import { BrowserRouter as Router, Route, Link } from "react-router-dom";

import BaseballTableWithData from "./BaseballTableWithData";

function AppRouter() {
  return (
    <Router>
      <div>
        <nav>
          <ul>
            <li>
              <Link to="/">Chris</Link>
            </li>
            <li>
              <Link to="/cork/">Cork</Link>
            </li>
          </ul>
        </nav>

        <Route path="/" exact render={() => <BaseballTableWithData />} />
        <Route
          path="/cork"
          render={() => <BaseballTableWithData league="cork" />}
        />
      </div>
    </Router>
  );
}

export default AppRouter;
