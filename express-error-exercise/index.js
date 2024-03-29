const tracer = require("dd-trace").init({
  service: "express-error-exercise",
  env: "sandbox",
  version: "1.0.0",
});
const express = require("express");
const app = express();
const port = 3001;

//This adds the callback function for the express plugin to validate http status codes that are < 400. For more inforamtion on this please refer to https://datadoghq.dev/dd-trace-js/interfaces/plugins.express.html#validatestatus
tracer.use("express", {
  validateStatus: (code) => code < 400,
});

app.post("/process_order_34e89a3f-b30c-4c1b-89c3-f94153d4d194", (req, res) => {
  res.status(200).send("Order: 34e89a3f-b30c-4c1b-89c3-f94153d4d194 has been processed.");
});


//PROPMT #1
//If a customer wants to observe the full error stack, they can impliment Custom Instrumentation to add the necessary error tags to their. Source: https://docs.datadoghq.com/tracing/trace_collection/custom_instrumentation/nodejs/?tab=errors#pagetitle
app.get("/test/1234567891245643758", (req, res) => {
  const span = tracer.scope().active()
  // span.setTag("resource_name", "127.33.678.1")
  span.setTag("test_tag", [123, 456])
  res.status(200).send("test");

  /*
throw new Error("There was an error!!!!") <---This works in this example

but for most customer's use, you would typically refer them to use:
  const span = tracer.scope().active()
  try {
    getIngredients()
  } catch (e) {
    span.setTag('error', e)
  }

*/
});

//validateStatus will mark 4xx http responses as errors
app.get("/error2", (req, res) => {
  res.status(400).send("Bad Request");
});

app.get("/error_tracking", (req, res) => {
  const span = tracer.scope().active();
  //If there are any errors occuring on the active span, the tag track_error true will ensure that the error is registered for APM error tracking as by default errors will only show if they are the service entry span. Source: https://docs.datadoghq.com/tracing/error_tracking/#pagetitle
  span.setTag("track_error", true);

  res.status(400).send("Bad Request");
});

app.listen(port, () => {
  console.log(`Example app listening on port ${port}`);
});
