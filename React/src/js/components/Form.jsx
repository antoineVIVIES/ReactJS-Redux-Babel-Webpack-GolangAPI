import React, { Component } from 'react';
import { connect } from 'react-redux';
import uuidv1 from 'uuid';
import addArticle from '../actions/index';

function mapDispatchToProps(dispatch) {
  return {
    addArticle: article => dispatch(addArticle(article))
  };
}
class ConnectedForm extends Component {
  constructor() {
    super();
    this.state = {
      title: '',
      name: '',
    };
    this.handleChange = this.handleChange.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
  }

  handleChange(event) {
    this.setState({ [event.target.id]: event.target.value });
  }

  handleSubmit(event) {
    event.preventDefault();
    const { title } = this.state;
    const { name } = this.state;
    if (name === '' || title === '') {
      alert('aeazeazeaze'); // eslint-disable-line
      return;
    }
    const id = uuidv1();
    this.props.addArticle({ title, id, name });
    this.setState({ title: '', name: '' });
  }

  render() {
    const { title } = this.state;
    const { name } = this.state;
    return (
      <form onSubmit={this.handleSubmit}>
        <div className="form-group">
          <label htmlFor="title">Title</label>
          <input
            type="text"
            className="form-control"
            id="title"
            value={title}
            onChange={this.handleChange}
          />
          <label htmlFor="name">Name</label>
          <input
            type="text"
            className="form-control"
            id="name"
            value={name}
            onChange={this.handleChange}
          />
        </div>
        <button type="submit" className="btn btn-success btn-lg">
          SAVE
        </button>
      </form>
    );
  }
}
const Form = connect(null, mapDispatchToProps)(ConnectedForm);
export default Form;
