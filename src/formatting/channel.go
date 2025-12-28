
package formatting

import . "TypeGo/core"



func ProcessChannel(format_data *FormatData, const_token Token) BlockData {
	
	format_data.AddToFunctionLog("ENTER ProcessChannel")
	
	var block_data BlockData
	format_data.Increment(); 
	var nextToken Token  = format_data.GetToken(); 
	if nextToken.Type == TokenType.NA {
		format_data.EndOfFileError(const_token); 
		format_data.AddToFunctionLog("ERROR ProcessChannel")
		return block_data; 
	}
	block_data = DeclarationLoop(format_data, nextToken); 
	
	SetNodeType( &block_data); 
	block_data.StartingToken = const_token; 
	format_data.AddToFunctionLog("EXIT ProcessChannel")
	return block_data; 
}

func SetNodeType(block_data *BlockData) {
	switch block_data.NodeType {
	case NodeType.Single_Declaration_With_Value, NodeType.Multiple_Declarations_With_Value, 
	NodeType.Multiple_Declarations_Same_Type_With_Value:
	
	block_data.NodeType = NodeType.Channel_Declaration_With_Value; 
	
	case NodeType.Single_Declaration_No_Value, NodeType.Multiple_Declarations_No_Value, 
	NodeType.Multiple_Declarations_Same_Type_No_Value:
	
	block_data.NodeType = NodeType.Channel_Declaration; 
	
	default:
	block_data.NodeType = NodeType.Channel_Declaration; 
	
	}
	
	
}
