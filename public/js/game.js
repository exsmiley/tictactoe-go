var Board = React.createClass({
  render: function() {
    return (
      <table>
      	<tr>
      		<Square />
      		<Square />
      		<Square />
      	</tr>
      	<tr>
      		<Square />
      		<Square />
      		<Square />
      	</tr>
      	<tr>
      		<Square />
      		<Square />
      		<Square />
      	</tr>
      </table>
    );
  }
});

var Square = React.createClass({
	render: function () {
		return (
			<td><img src="img/o.png"/></td>
		);
	}
});

ReactDOM.render(
  <Board />,
  document.getElementById('content')
);