var pics = {"x": "img/x.png", "o": "img/o.png", "": ""}

class Board extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            board: [['','',''], ['','',''], ['','','']],
        };
        this.handleClick = this.handleClick.bind(this);
    }
    handleClick(x,y) {
        // update based on click
        var board = this.state.board
        board[x][y] = 'x'

        // get AI move
        var xhttp = new XMLHttpRequest();
        xhttp.open("POST", "/move", true);
        xhttp.setRequestHeader("Content-Type", "application/json");
        
        var data = JSON.stringify({"board": board})
        xhttp.send(data);

        xhttp.onreadystatechange = function() {
            if (xhttp.readyState == 4 && xhttp.status == 200) {
                console.log(xhttp.responseText);
            }
        };

        const x2 = Math.floor((Math.random() * 3) + 0);
        const y2 = Math.floor((Math.random() * 3) + 0);
        board[x2][y2] = 'o'
        this.setState({board: board});
    }
    render() {
        return (
            <table>
            <tbody>
            <tr>
            <Square board={this.state.board} handleClick={this.handleClick} x="0" y="0"/>
            <Square board={this.state.board} handleClick={this.handleClick} x="0" y="1"/>
            <Square board={this.state.board} handleClick={this.handleClick} x="0" y="2"/>
            </tr>
            <tr>
            <Square board={this.state.board} handleClick={this.handleClick} x="1" y="0"/>
            <Square board={this.state.board} handleClick={this.handleClick} x="1" y="1"/>
            <Square board={this.state.board} handleClick={this.handleClick} x="1" y="2"/>
            </tr>
            <tr>
            <Square board={this.state.board} handleClick={this.handleClick} x="2" y="0"/>
            <Square board={this.state.board} handleClick={this.handleClick} x="2" y="1"/>
            <Square board={this.state.board} handleClick={this.handleClick} x="2" y="2"/>
            </tr>
            </tbody>
            </table>
            );  
    }
}

class Square extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            value: pics[props.board[props.x][props.y]],
            x: props.x,
            y: props.y
        };
        this.handleClick = this.handleClick.bind(this);
    }
    handleClick() {
    this.props.handleClick(this.state.x, this.state.y)
    }
    componentWillReceiveProps(props) {
        let pic = pics[props.board[this.state.x][this.state.y]]
        this.setState({value: pic});
    }
    render() {
        return (
            <td onClick={this.handleClick}>
            <img src={this.state.value}/>
            </td>
            );
    }
}

ReactDOM.render(
    <Board />,
    document.getElementById('content')
    );