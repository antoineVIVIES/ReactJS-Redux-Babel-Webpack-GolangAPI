import React from 'react';
import { connect } from 'react-redux';

const mapStateToProps = (state) => {
  return { articles: state.articles };
};
const ConnectedList = ({ articles }) => (
  <ul className="list-group list-group-flush">
    {articles.map(el => (
      <li className="list-group-item" key={el.id}>
        <ul>
          {el.title}
        </ul>
        <ul>
          {el.name}
        </ul>
      </li>
    ))}
  </ul>
);
const List = connect(mapStateToProps)(ConnectedList);
export default List;
