/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface ControllersRoomuse
 */
export interface ControllersRoomuse {
    /**
     * 
     * @type {string}
     * @memberof ControllersRoomuse
     */
    contact?: string;
    /**
     * 
     * @type {string}
     * @memberof ControllersRoomuse
     */
    inTime?: string;
    /**
     * 
     * @type {string}
     * @memberof ControllersRoomuse
     */
    note?: string;
    /**
     * 
     * @type {string}
     * @memberof ControllersRoomuse
     */
    outTime?: string;
    /**
     * 
     * @type {number}
     * @memberof ControllersRoomuse
     */
    people?: number;
    /**
     * 
     * @type {number}
     * @memberof ControllersRoomuse
     */
    purposeID?: number;
    /**
     * 
     * @type {number}
     * @memberof ControllersRoomuse
     */
    roomID?: number;
    /**
     * 
     * @type {number}
     * @memberof ControllersRoomuse
     */
    userID?: number;
}

export function ControllersRoomuseFromJSON(json: any): ControllersRoomuse {
    return ControllersRoomuseFromJSONTyped(json, false);
}

export function ControllersRoomuseFromJSONTyped(json: any, ignoreDiscriminator: boolean): ControllersRoomuse {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'contact': !exists(json, 'contact') ? undefined : json['contact'],
        'inTime': !exists(json, 'inTime') ? undefined : json['inTime'],
        'note': !exists(json, 'note') ? undefined : json['note'],
        'outTime': !exists(json, 'outTime') ? undefined : json['outTime'],
        'people': !exists(json, 'people') ? undefined : json['people'],
        'purposeID': !exists(json, 'purposeID') ? undefined : json['purposeID'],
        'roomID': !exists(json, 'roomID') ? undefined : json['roomID'],
        'userID': !exists(json, 'userID') ? undefined : json['userID'],
    };
}

export function ControllersRoomuseToJSON(value?: ControllersRoomuse | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'contact': value.contact,
        'inTime': value.inTime,
        'note': value.note,
        'outTime': value.outTime,
        'people': value.people,
        'purposeID': value.purposeID,
        'roomID': value.roomID,
        'userID': value.userID,
    };
}


