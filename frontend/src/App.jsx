import { Router, Route } from "@solidjs/router";

import Home from "./home/HomePage.jsx";
import About from "./about/AboutPage.jsx";
import Quote from "./quote/QuotePage.jsx";


export default function App() {
  return (
    <Router>
      <Route path="/" component={Home} />
      <Route path="/about" component={About} />
      <Route path="/quote" component={Quote} />
    </Router>
  );
}
