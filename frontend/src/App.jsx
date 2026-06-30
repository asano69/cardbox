import { Router, Route } from "@solidjs/router";
import Home from "./HomePage.jsx";
import About from "./AboutPage.jsx";
import QuoteInputPage from "./quote/QuoteInputPage.jsx"

export default function App() {
  return (
    <Router>
      <Route path="/" component={Home} />
      <Route path="/about" component={About} />
      <Route path="/new" component={QuoteInputPage} />
    </Router>
  );
}
 


