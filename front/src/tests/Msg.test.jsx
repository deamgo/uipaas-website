// index.test.js
import { render, act, getByTitle } from "@testing-library/react";
import "@testing-library/jest-dom";
import { MessageContainer } from "../components/Msg";
import $message from "../components/Msg";
import { MessageApi } from "../components/Msg/config";
import Msg from "../components/Msg/Msg";

test("renders messages correctly", () => {
  const { getByText } = render(<MessageContainer />);

  //   $message.info("This is an info message");

  //   expect(getByTitle("This is an info message")).toBeInTheDocument();
  expect($message.info("This is an info message"));
});

test("renders messages correctly", () => {
  const { getByText } = render(<MessageContainer />);
});

test("handles multiple messages correctly", () => {
  const { getByText } = render(<MessageContainer />);

  expect($message.success("Success message"));
  expect($message.warning("Warning message"));
  expect($message.error("Error message"));
});

test("message api", () => {
expect(MessageApi.info())
expect(MessageApi.success())
expect(MessageApi.error())
expect(MessageApi.warning())
})

describe('Msg', () => {
    it('renders correctly', () => {
      const { container } = render(<Msg type="success" text="This is a success message." />);
      expect(container.firstChild).toMatchSnapshot();
    });
  });