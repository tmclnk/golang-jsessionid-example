# JSESSIONID Login Example

This repo has a class called `Client` that wraps an `http.Client` with with a baked in cookie jar.

Pay special attention to the Login() method's handling of query params (I didn't find this to be obvious).