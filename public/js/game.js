var pics = {"x": "img/x.png", "o": "img/o.png", "": ""}

class Board extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            board: [['','',''], ['','',''], ['','','']],
            error: "",
            winner: "",
            playerStart: true,
            over: false
        };
        this.handleMove = this.handleMove.bind(this);
        this.newGame = this.newGame.bind(this);
    }
    handleMove(x,y) {
        if (!this.state.over) {
            var board = this.state.board

            // try to move on a spot that has been taken already
            if (board[x][y] != "") {
                this.setState({error: board[x][y].toUpperCase() + " is already there"})
                return
            }
            else {
                this.setState({error: ""})
            }

            // update based on click
            board[x][y] = 'x'

            // get AI move
            var xhttp = new XMLHttpRequest();
            xhttp.open("POST", "/move", true);
            xhttp.setRequestHeader("Content-Type", "application/json");

            var data = JSON.stringify({"board": board, "winner": ""})
            xhttp.send(data);

            // super hacky solution to solve a scoping problem... I don't like it
            var obj = this

            xhttp.onreadystatechange = function() {
                if (xhttp.readyState == 4 && xhttp.status == 200) {
                    const response = JSON.parse(xhttp.responseText)
                    const board = response['Board'];
                    obj.setState({board:board});

                    const winner = response['Winner'];
                    if (winner != "") {
                        obj.setState({error: winner.toUpperCase() + " won the game!"})
                        obj.setState({winner: winner})
                        obj.setState({over: true})
                        if (winner == "cat") {
                            const playerStart = !obj.state.playerStart
                            obj.setState({cat: cat}) 
                        }
                    }
                }
            }
        }
    }
    newGame() {
        this.setState({over: false, error: ""})
        if (this.state.winner == "" || this.state.winner == "o") {
            this.setState({
                board: [['','',''], ['','',''], ['','','']]})
        }
        else if (this.state.cat == 1 && this.state.playerStart) {
            this.setState({
                board: [['','',''], ['','',''], ['','','']],
                error: ""})
        }
        else {
            console.log(this.state.winner)
            var board = [['','',''], ['','',''], ['','','']]
            // get AI move
            var xhttp = new XMLHttpRequest();
            xhttp.open("POST", "/move", true);
            xhttp.setRequestHeader("Content-Type", "application/json");

            var data = JSON.stringify({"board": board, "over": false})
            xhttp.send(data);

            // super hacky solution to solve a scoping problem... I don't like it
            var obj = this

            xhttp.onreadystatechange = function() {
                if (xhttp.readyState == 4 && xhttp.status == 200) {
                    const response = JSON.parse(xhttp.responseText)
                    const board = response['Board'];
                    obj.setState({board:board});
                }
            };
        }
    }
    render() {
        return (
            <div>
                <table>
                    <tbody>
                        <tr>
                            <Square board={this.state.board} handleClick={this.handleMove} x="0" y="0"/>
                            <Square board={this.state.board} handleClick={this.handleMove} x="0" y="1"/>
                            <Square board={this.state.board} handleClick={this.handleMove} x="0" y="2"/>
                        </tr>
                        <tr>
                            <Square board={this.state.board} handleClick={this.handleMove} x="1" y="0"/>
                            <Square board={this.state.board} handleClick={this.handleMove} x="1" y="1"/>
                            <Square board={this.state.board} handleClick={this.handleMove} x="1" y="2"/>
                        </tr>
                        <tr>
                            <Square board={this.state.board} handleClick={this.handleMove} x="2" y="0"/>
                            <Square board={this.state.board} handleClick={this.handleMove} x="2" y="1"/>
                            <Square board={this.state.board} handleClick={this.handleMove} x="2" y="2"/>
                        </tr>
                    </tbody>
                </table>
                <h1>{this.state.error}</h1>
                <button onClick={this.newGame}>New Game</button>
            </div>
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