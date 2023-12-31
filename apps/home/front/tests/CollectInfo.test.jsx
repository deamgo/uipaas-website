import React from "react";
import { vi } from 'vitest';
import { render, screen, fireEvent } from "@testing-library/react";
import "@testing-library/jest-dom";
import { CollectInfo } from "../src/components/CollectInfo";
import $message from "../src/components/Msg";


describe("CollectInfo Component", () => {
  it("renders the component", () => {
    render(<CollectInfo />);
    expect(screen.getByText("Hello, future partners!")).toBeInTheDocument();
    expect(screen.getByText("Company")).toBeInTheDocument();
  });

  it("submits the form with valid data", () => {
    render(<CollectInfo />);
    fireEvent.change(screen.getByRole('textbox', { name: 'Company *' }), {
      target: { value: "Test Company" },
    });
    fireEvent.change(screen.getByRole('textbox', { name: 'Team Size *' }), {
      target: { value: "Test Size" },
    });
    fireEvent.change(screen.getByRole('textbox', { name: 'Name *' }), {
      target: { value: "Test Name" },
    });
    fireEvent.change(screen.getByRole('textbox', { name: 'Business email *' }), {
      target: { value: "test@example.com" },
    });
    fireEvent.change(screen.getByRole('textbox', { name: 'Description of Requirements *' }), {
      target: { value: "Test Description" },
    });

    fireEvent.submit(screen.getByRole("button", { name: "Submit" }));

  });

  it("displays error message on form submission with invalid data", () => {
    vi.spyOn($message, 'error')
    render(<CollectInfo />);
    fireEvent.submit(screen.getByRole("button", { name: "Submit" }));

    expect($message.error).toHaveBeenCalledWith("companyname cant be empty");
  });
});
