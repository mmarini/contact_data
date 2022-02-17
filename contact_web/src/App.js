import React from "react";
import ReactDOM from "react-dom";

import './App.css';

export default class ContactForm extends React.Component {
  constructor() {
    super();
    this.state = {
      full_name: "",
      email: "",
      phone_numbers: [{ number: "" }],
      message: ""
    };
  }

  handleNameChange = evt => {
    this.setState({ full_name: evt.target.value });
  };

  handleEmailChange = evt => {
    this.setState({ email: evt.target.value });
  };

  handlePhoneNumberChange = idx => evt => {
    const newPhoneNumbers = this.state.phone_numbers.map((phone_number, sidx) => {
      if (idx !== sidx) return phone_number;
      return { ...phone_number, number: evt.target.value };
    });

    this.setState({ phone_numbers: newPhoneNumbers });
  };

  handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const { full_name, email, phone_numbers } = this.state;

      var numbers = phone_numbers.map( function(pn) {
          return pn.number;
        }
      ) 

      let res = await fetch("http://localhost:8080/contacts", {
        method: "POST",
        body: JSON.stringify({
          full_name: full_name,
          email: email,
          phone_numbers: numbers
        }),
      });
      let resJson = await res.json();
      if (res.status === 201) {
        this.setState({ full_name: "" });
        this.setState({ email: "" });
        this.setState({ phone_numbers: [{ number: "" }] });
        this.setState({ message: "Contact created successfully" })
      } else {
        this.setState({ message: "An error occured" })
      }
    } catch (err) {
      console.log(err);
    }
  };

  handleAddPhoneNumber = () => {
    this.setState({
      phone_numbers: this.state.phone_numbers.concat([ "" ])
    });
  };

  handleRemovePhoneNumber = idx => () => {
    this.setState({
      phone_numbers: this.state.phone_numbers.filter((s, sidx) => idx !== sidx)
    });
  };

  render() {
    return (
      <div className="wrapper">
        <div className="message">{this.state.message ? <p>{this.state.message}</p> : null}</div>
        <h1>Contact Data</h1>
        <form onSubmit={this.handleSubmit}>
          <fieldset>
            <label>
              <p>Full Name</p>
              <input
                type="text"
                value={this.state.name}
                onChange={this.handleNameChange}
              />
            </label>

            <label>
              <p>Email</p>
              <input
                type="text"
                value={this.state.email}
                onChange={this.handleEmailChange}
              />
            </label>

            <h4>Phone Numbers</h4>

            {this.state.phone_numbers.map((phone_number, idx) => (
              <div className="phone_number">
                <input
                  type="text"
                  value={phone_number.number}
                  onChange={this.handlePhoneNumberChange(idx)}
                />
                <button
                  type="button"
                  onClick={this.handleRemovePhoneNumber(idx)}
                  className="small"
                >
                  -
                </button>
              </div>
            ))}

            <button
              type="button"
              onClick={this.handleAddPhoneNumber}
              className="small"
            >
              Add Phone Number
            </button>
            <br/>
            <button>Submit</button>
          </fieldset>
        </form>
      </div>
    );
  }
}
