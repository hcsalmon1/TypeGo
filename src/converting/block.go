
package converting

import . "TypeGo/core"



func ProcessBlockData(convertData *ConvertData, blockData *BlockData, nestCount int, isFirst bool) {
	
	var nodeType IntNodeType  = blockData.NodeType; 
	
	var wasIfOrElseIf bool  = 
	
	convertData.LastNodeType == NodeType.If_Statement || 
	
	convertData.LastNodeType == NodeType.Else_Statement; 
	
	if wasIfOrElseIf {
		if nodeType != NodeType.Else_Statement {
			//convertData.NewLineWithTabs(nestCount);
			
		}
		convertData.LastNodeType = NodeType.Invalid; 
		
	}
	
	switch nodeType {
	
	case NodeType.Invalid:
	convertData.ConvertResult = ConvertResult.Invalid_Node_Type; 
	convertData.ErrorDetail = "invalid node type in block data"; 
	convertData.ErrorToken = blockData.StartingToken; 
	return; 
	
	case NodeType.Single_Declaration_With_Value:
	ProcessSingleDeclarationWithValue(convertData, blockData, nestCount); 
	
	case NodeType.Single_Declaration_No_Value:
	ProcessSingleDeclarationNoValue(convertData, blockData, nestCount); 
	
	case NodeType.Multiple_Declarations_No_Value:
	ProcessMultipleDeclarationNoValue(convertData, blockData, nestCount); 
	
	case NodeType.Multiple_Declarations_With_Value, NodeType.Multiple_Declarations_Same_Type_No_Value:
	convertData.ConvertResult = ConvertResult.Unsupported_Type; 
	convertData.ErrorDetail = "unsupported node type: {nodeType}"; 
	convertData.ErrorToken = blockData.StartingToken; 
	
	case NodeType.Multiple_Declarations_Same_Type_With_Value, NodeType.Multiple_Declarations_One_Type_One_Set_Value:
	ProcessMultipleDeclarationWithSetValue(convertData, blockData, nestCount); 
	
	case NodeType.If_Statement:
	PrintIfStatement(convertData, blockData, nestCount); 
	
	case NodeType.Else_Statement:
	PrintElseStatement(convertData, blockData, nestCount); 
	
	case NodeType.For_Loop:
	PrintFor(convertData, blockData, nestCount); 
	
	case NodeType.Multi_Line_Import, NodeType.Single_Import, NodeType.Single_Import_With_Alias, NodeType.Package:
	PrintTokensNewLine(convertData, blockData, nestCount, true); 
	
	case NodeType.Other:
	PrintTokens(convertData, blockData, nestCount); 
	
	case NodeType.Switch:
	ConvertSwitch(convertData, blockData, nestCount, false); 
	
	case NodeType.Struct_Declaration:
	ProcessStruct(convertData, blockData, nestCount); 
	
	case NodeType.Enum_Declaration:
	ProcessEnum(convertData, blockData, nestCount); 
	
	case NodeType.Enum_Struct_Declaration:
	ProcessEnumstruct(convertData, blockData, nestCount); 
	
	case NodeType.Enum_Struct_Declaration_With_Alias:
	ProcessEnumstructWithAlias(convertData, blockData, nestCount); 
	
	case NodeType.Append:
	PrintAppend(convertData, blockData, nestCount); 
	
	case NodeType.Constant_Global_Variable:
	convertData.AppendString("const "); 
	PrintTokens(convertData, blockData, nestCount); 
	
	case NodeType.Constant_Global_Variable_With_Type:
	PrintConstant(convertData, blockData, nestCount); 
	
	case NodeType.Comment:
	PrintComment(convertData, blockData, nestCount); 
	
	case NodeType.Channel_Declaration:
	ProcessChannelDeclarationNoValue(convertData, blockData, nestCount); 
	
	case NodeType.Channel_Declaration_With_Value:
	ProcessChannelDeclarationWithValue(convertData, blockData, nestCount); 
	
	case NodeType.Interface_Declaration:
	ProcessInterfaceDeclaration(convertData, blockData, nestCount); 
	
	case NodeType.NewLine:
	convertData.NewLineWithTabs(); 
	
	case NodeType.Err_Return:
	Process(convertData, blockData, nestCount); 
	
	case NodeType.Err_Check:
	PrintErrorCheckStatement(convertData, blockData, nestCount); 
	
	case NodeType.Return:
	PrintTokens(convertData, blockData, nestCount); 
	
	case NodeType.Break:
	convertData.AppendString("break"); 
	
	default:
	convertData.ConvertResult = ConvertResult.Unexpected_Type; 
	convertData.ErrorDetail = "unexpected node type '{nodeType}' in block data"; 
	convertData.ErrorToken = blockData.StartingToken; 
	break; 
	
	}
	
	
	AddSpaceAfterBlock(convertData, blockData.NodeType, nestCount); 
	
}

func ProcessBlock(convertData *ConvertData, codeBlock *CodeBlock, nestCount int, isGlobal bool) {
	
	convertData.SetErrorFunction("ProcessBlock")
	
	var blockDataList []BlockData  = codeBlock.BlockDataList; 
	if len(blockDataList) == 0 {
		return; 
	}
	convertData.NewLineWithTabs(); 
	
	var isFirst bool  = true; 
	
	for i := 0; i < len(blockDataList); i++ {
	
		var blockData BlockData  = blockDataList[i]; 
		if blockData.NodeType == NodeType.NewLine && i == 0 {
			continue 
			
			
		}
		ProcessBlockData(convertData, &blockData, nestCount, isFirst); 
		if isGlobal == false {
			isFirst = false; 
			
		}
		if convertData.IsError() {
			return 
		}
		
	}
	
	
}
